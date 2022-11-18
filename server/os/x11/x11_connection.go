package x11

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xinerama"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

const (
	xfixesMajorVersion = 5
	xfixesMinorVersion = 0
)

type connection struct {
	impl         *xgb.Conn
	atoms        *atomStorage
	setupInfo    *xproto.SetupInfo
	eventsCh     chan<- interface{}
	errorCtx     *module.ErrorCtx
	moduleLogger *zap.Logger
	eventLoopWG  sync.WaitGroup
	quit         int32
}

func newConnection() *connection {
	return &connection{
		impl:         nil,
		atoms:        nil,
		setupInfo:    nil,
		eventsCh:     nil,
		errorCtx:     nil,
		moduleLogger: nil,
		eventLoopWG:  sync.WaitGroup{},
		quit:         0,
	}
}

func (c *connection) init(
	display string, atoms *atomStorage, eventsCh chan<- interface{}, errorCtx *module.ErrorCtx, moduleLogger *zap.Logger) bool {

	c.atoms = atoms
	c.eventsCh = eventsCh
	c.errorCtx = errorCtx
	c.moduleLogger = moduleLogger

	if !c.createImpl(display) {
		c.stop()
	}

	return true
}

func (c *connection) start() {
	go func() {
		c.eventLoopWG.Add(1)
		defer c.eventLoopWG.Done()
		if err := c.eventLoop(); err != nil {
			c.errorCtx.SendError(err)
		}
	}()
}

func (c *connection) eventLoop() error {
	var ev xgb.Event
	var err xgb.Error
	var resultErr error

	eventsCh := c.eventsCh
	for atomic.LoadInt32(&c.quit) == 0 {
		if ev, err = c.impl.WaitForEvent(); err != nil {
			c.moduleLogger.Debug("x11 event error",
				zap.Uint32("BadId", err.BadId()),
				zap.Uint16("SequenceId", err.SequenceId()),
				zap.Error(err),
			)

			continue
		}

		if ev == nil {
			resultErr = errors.New("Expected an x11 event but got nil")
			c.moduleLogger.Error("Expected an x11 event but got nil")
			break
		}

		switch event := ev.(type) {
		case xproto.MappingNotifyEvent:
			eventsCh <- event
		case xproto.KeyReleaseEvent:
			eventsCh <- event
		case xfixes.SelectionNotifyEvent:
			eventsCh <- event
		case xproto.SelectionNotifyEvent:
			eventsCh <- event
		case xproto.PropertyNotifyEvent:
			eventsCh <- event
		case xproto.SelectionRequestEvent:
			eventsCh <- event
		case xproto.SelectionClearEvent:
			eventsCh <- event
		}
	}

	return resultErr
}

func (c *connection) stop() {
	if c.impl != nil {
		atomic.StoreInt32(&c.quit, 0)
		c.eventLoopWG.Wait()
		c.impl.Close()
		c.impl = nil
	}
}

func (c *connection) createImpl(display string) bool {
	var err error
	if c.impl, err = xgb.NewConnDisplay(display); err != nil {
		c.moduleLogger.Warn("Failed connect to x11 server", zap.Error(err))
		return false
	}

	if err = xinerama.Init(c.impl); err != nil {
		c.moduleLogger.Warn("Failed init xinerama extension", zap.Error(err))
		return false
	}

	// https://www.x.org/releases/X11R7.7/doc/fixesproto/fixesproto.txt
	if err := xfixes.Init(c.impl); err != nil {
		c.moduleLogger.Warn("Failed init xfixes extension", zap.Error(err))
		return false
	}

	xfixesVersion, err := xfixes.QueryVersion(c.impl, xfixesMajorVersion, xfixesMinorVersion).Reply()
	if err != nil {
		c.moduleLogger.Warn("Failed get xfixes extension version", zap.Error(err))
		return false
	}

	if (xfixesVersion.MajorVersion < xfixesMajorVersion) ||
		(xfixesVersion.MajorVersion == xfixesMajorVersion && xfixesVersion.MinorVersion < xfixesMinorVersion) {
		c.moduleLogger.Warn("Wrong xfixes extension version",
			zap.Uint32("expectedMajorVersion", xfixesMajorVersion),
			zap.Uint32("expectedMinorVersion", xfixesMinorVersion),
			zap.Uint32("actualMajorVersion", xfixesVersion.MajorVersion),
			zap.Uint32("actualMinorVersion", xfixesVersion.MinorVersion),
		)

		return false
	}

	c.setupInfo = xproto.Setup(c.impl)

	return true
}

func (c *connection) getSetupInfo() *xproto.SetupInfo {
	return c.setupInfo
}

func (c *connection) getDefaultScreen() *xproto.ScreenInfo {
	return c.setupInfo.DefaultScreen(c.impl)
}

func (c *connection) sendEvent(windowID xproto.Window, eventMask int, event string, fields ...zap.Field) bool {
	if err := xproto.SendEventChecked(c.impl, false, windowID, uint32(eventMask), event).Check(); err != nil {
		c.moduleLogger.Info("Failed create x11 graphics context",
			append(fields,
				zap.Error(err),
			)...)
		return false
	}

	return true
}

func (c *connection) getKeyboardMapping(fields ...zap.Field) (*xproto.GetKeyboardMappingReply, bool) {
	firstKeycode := c.setupInfo.MinKeycode
	count := byte(c.setupInfo.MaxKeycode - firstKeycode + 1)
	if keymap, err := xproto.GetKeyboardMapping(c.impl, firstKeycode, count).Reply(); err != nil {
		c.moduleLogger.Warn("Failed get keyboard mapping",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	} else {
		return keymap, true
	}
}

func (c *connection) getModifierMapping(fields ...zap.Field) (*xproto.GetModifierMappingReply, bool) {
	if modmap, err := xproto.GetModifierMapping(c.impl).Reply(); err != nil {
		c.moduleLogger.Warn("Failed get modifier mapping",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	} else {
		return modmap, true
	}
}

func (c *connection) createAtom(name atomName, fields ...zap.Field) (xproto.Atom, bool) {
	r, err := xproto.InternAtom(c.impl, false, uint16(len(name)), string(name)).Reply()
	if err != nil {
		c.moduleLogger.Warn("Failed create x11 atom",
			append(fields,
				name.ZapField(),
				zap.Error(err),
			)...)
		return xproto.AtomNone, false
	}
	if r == nil {
		c.moduleLogger.Warn("Failed create x11 atom", name.ZapField())
		return xproto.AtomNone, false
	}

	return r.Atom, true
}

func (c *connection) getAtom(id xproto.Atom, fields ...zap.Field) (atomName, bool) {
	reply, err := xproto.GetAtomName(c.impl, id).Reply()
	if err != nil {
		c.moduleLogger.Warn("Failed get x11 atom name",
			append(fields,
				zap.Uint32("AtomID", uint32(id)),
				zap.Error(err),
			)...)
		return "", false
	}

	if reply == nil {
		c.moduleLogger.Warn("Failed get x11 atom name",
			append(fields,
				zap.Uint32("AtomID", uint32(id)),
			)...)
		return "", false
	}

	return atomName(reply.Name), true
}

func (c *connection) newWindow(windowID xproto.Window) *window {
	return newWindow(c, windowID)
}

func (c *connection) createWindow(rootWindowID xproto.Window, screen *xproto.ScreenInfo, fields ...zap.Field) (*window, bool) {
	var err error

	// Create a general purpose graphics context
	var gc xproto.Gcontext
	if gc, err = xproto.NewGcontextId(c.impl); err != nil {
		c.moduleLogger.Info("Failed create x11 graphics context id",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	}

	if err = xproto.CreateGCChecked(
		c.impl, gc, xproto.Drawable(rootWindowID), xproto.GcForeground, []uint32{screen.WhitePixel}).Check(); err != nil {
		c.moduleLogger.Info("Failed create x11 graphics context",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	}

	// Create a window
	var windowID xproto.Window
	if windowID, err = xproto.NewWindowId(c.impl); err != nil {
		c.moduleLogger.Info("Failed create x11 window id",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	}

	err = xproto.CreateWindowChecked(
		c.impl, screen.RootDepth, windowID, rootWindowID,
		-1000, -1000, 1, 1, 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwEventMask|xproto.CwOverrideRedirect,
		[]uint32{1, xproto.EventMaskPropertyChange}).Check()
	if err != nil {
		c.moduleLogger.Info("Failed create x11 window",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	}

	err = xproto.MapWindowChecked(c.impl, windowID).Check()
	if err != nil {
		c.moduleLogger.Info("Failed map x11 window",
			append(fields,
				zap.Error(err),
			)...)
		return nil, false
	}

	return c.newWindow(windowID), true
}

func (c *connection) grabKey(windowID xproto.Window, modifiers uint16, keycode xproto.Keycode, fields ...zap.Field) global.Error {
	if err := xproto.GrabKeyChecked(
		c.impl, true, windowID, modifiers, keycode, xproto.GrabModeAsync, xproto.GrabModeAsync).Check(); err != nil {
		switch err.(type) {
		case xproto.AccessError:
			accessErr := errors.New("keyboard shortcut is already taken by another application")
			c.moduleLogger.Debug("Failed call x11 grab key",
				append(fields,
					zap.Uint8("X11Keycode", uint8(keycode)),
					zap.Uint16("X11Modifiers", modifiers),
					zap.Error(accessErr),
				)...)
			return global.HotkeyUsesByExternalApp
		default:
			c.moduleLogger.Warn("Failed call x11 grab key",
				append(fields,
					zap.Uint8("X11Keycode", uint8(keycode)),
					zap.Uint16("X11Modifiers", modifiers),
					zap.Error(err),
				)...)
			return global.HotkeyBindError
		}
	}

	return global.Success
}

func (c *connection) ungrabKey(windowID xproto.Window, modifiers uint16, keycode xproto.Keycode, fields ...zap.Field) bool {
	if err := xproto.UngrabKeyChecked(c.impl, keycode, windowID, modifiers).Check(); err != nil {
		c.moduleLogger.Warn("Failed call x11 ungrab key",
			append(fields,
				zap.Uint8("X11Keycode", uint8(keycode)),
				zap.Uint16("X11Modifiers", modifiers),
				zap.Error(err),
			)...)
		return false
	}

	return true
}

func (c *connection) subscribeToSelectionChanges(windowID xproto.Window, selection xproto.Atom, fields ...zap.Field) bool {
	if err := xfixes.SelectSelectionInputChecked(
		c.impl, windowID, selection, xfixes.SelectionEventMaskSetSelectionOwner).Check(); err != nil {
		c.moduleLogger.Info("Failed subscribe to selection buffer changes",
			append(fields,
				c.atoms.getZapFieldPrefix("Selection", selection),
				zap.Error(err),
			)...)
		return false
	}

	return true
}

func (c *connection) readProperty(
	windowID xproto.Window, property xproto.Atom, fields ...zap.Field) ([]byte, readPropertyResult) {

	deleteProperty := true
	result := make([]byte, 0, 1024)
	chunkSize := uint32(1024 * 1024) // 1 MB
	for {
		// convert size of bytes to size of longs
		longOffset := uint32(len(result)) / 4
		longLength := chunkSize / 4

		reply, err := xproto.GetProperty(
			c.impl, deleteProperty, windowID, property, xproto.AtomAny, longOffset, longLength).Reply()
		if err != nil {
			c.moduleLogger.Info("Failed call x11 get property",
				append(fields,
					c.atoms.getZapFieldPrefix("Property", property),
					zap.Error(err),
				)...)
			return []byte{}, rpFailed
		}

		if reply == nil {
			c.moduleLogger.Info("Failed call x11 get property",
				append(fields,
					c.atoms.getZapFieldPrefix("Property", property),
				)...)
			return []byte{}, rpFailed
		}

		if reply.Type == c.atoms.atomIncr {
			return []byte{}, rpIncremental
		}

		if len(reply.Value) > 0 {
			result = append(result, reply.Value...)
		}

		if reply.BytesAfter == 0 {
			break
		}
	}

	return result, rpSuccess
}

func (c *connection) writeProperty(
	windowID xproto.Window, property xproto.Atom, target xproto.Atom, itemLen byte, data []byte, fields ...zap.Field) bool {

	format := itemLen * 8
	dataLen := uint32(len(data) / int(itemLen))
	err := xproto.ChangePropertyChecked(
		c.impl, xproto.PropModeReplace, windowID, property, target, format, dataLen, data).Check()
	if err != nil {
		c.moduleLogger.Info("Failed call x11 change property",
			append(fields,
				c.atoms.getZapFieldPrefix("Property", property),
				zap.Error(err),
			)...)
		return false
	}

	return true
}

func (c *connection) getClipboardOwner(selection xproto.Atom, fields ...zap.Field) (xproto.Window, bool) {
	reply, err := xproto.GetSelectionOwner(c.impl, selection).Reply()
	if err != nil {
		c.moduleLogger.Info("Failed call x11 get clipboard owner",
			append(fields,
				c.atoms.getZapFieldPrefix("Selection", selection),
				zap.Error(err),
			)...)
		return 0, false
	}

	if reply.Owner == xproto.AtomNone {
		return 0, false
	}

	return reply.Owner, true
}

func (c *connection) setSelectionOwner(windowID xproto.Window, selection xproto.Atom, fields ...zap.Field) bool {
	err := xproto.SetSelectionOwnerChecked(c.impl, windowID, selection, xproto.TimeCurrentTime).Check()
	if err != nil {
		c.moduleLogger.Info("Failed call x11 set selection owner",
			append(fields,
				c.atoms.getZapFieldPrefix("Selection", selection),
				zap.Error(err),
			)...)
		return false
	}

	return true
}

func (c *connection) convertSelection(
	windowID xproto.Window, selection xproto.Atom, target xproto.Atom, property xproto.Atom, fields ...zap.Field) bool {

	err := xproto.ConvertSelectionChecked(
		c.impl, windowID, selection, target, property, xproto.TimeCurrentTime).Check()

	if err != nil {
		c.moduleLogger.Warn("Failed call x11 convert selection",
			append(fields,
				c.atoms.getZapFieldPrefix("Selection", selection),
				c.atoms.getZapFieldPrefix("Target", target),
				c.atoms.getZapFieldPrefix("Property", property),
				zap.Error(err),
			)...)
		return false
	}

	return true
}
