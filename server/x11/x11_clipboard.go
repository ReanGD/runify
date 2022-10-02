package x11

import (
	"errors"
	"fmt"

	"github.com/ReanGD/runify/server/config"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/jezek/xgbutil"
	"github.com/jezek/xgbutil/xevent"
	"go.uber.org/zap"
)

const (
	MB uint32 = 1024 * 1024
)

type x11Clipboard struct {
	xConnection      *xgbutil.XUtil
	atomTargets      xproto.Atom
	atomTimestamp    xproto.Atom
	atomUTF8String   xproto.Atom
	atomPrimarySel   xproto.Atom
	atomClipboardSel xproto.Atom
	ownClipboardData map[xproto.Atom]Mimes

	moduleLogger *zap.Logger
}

func newX11Clipboard() *x11Clipboard {
	return &x11Clipboard{
		xConnection:      nil,
		atomTargets:      0,
		atomTimestamp:    0,
		atomUTF8String:   0,
		atomPrimarySel:   0,
		atomClipboardSel: 0,
		ownClipboardData: make(map[xproto.Atom]Mimes),
		moduleLogger:     nil,
	}
}

func (h *x11Clipboard) getWindow() xproto.Window {
	return h.xConnection.Dummy()
}

func (h *x11Clipboard) getConnection() *xgb.Conn {
	return h.xConnection.Conn()
}

func (h *x11Clipboard) createAtom(name string) (xproto.Atom, error) {
	r, err := xproto.InternAtom(h.getConnection(), false, uint16(len(name)), name).Reply()
	if err != nil {
		h.moduleLogger.Error("Failed create x11 atom", zap.String("name", name), zap.Error(err))
		return 0, errors.New("failed init x11Clipboard module")
	}
	if r == nil {
		h.moduleLogger.Error("Failed create x11 atom", zap.String("name", name))
		return 0, errors.New("failed init x11Clipboard module")
	}

	return r.Atom, nil
}

func (h *x11Clipboard) onInit(cfg *config.Config, xConnection *xgbutil.XUtil, moduleLogger *zap.Logger) error {
	h.xConnection = xConnection
	h.moduleLogger = moduleLogger

	var err error

	if h.atomTargets, err = h.createAtom("TARGETS"); err != nil {
		return err
	}

	if h.atomTimestamp, err = h.createAtom("TIMESTAMP"); err != nil {
		return err
	}

	if h.atomUTF8String, err = h.createAtom("UTF8_STRING"); err != nil {
		return err
	}

	if h.atomPrimarySel, err = h.createAtom("PRIMARY"); err != nil {
		return err
	}

	if h.atomClipboardSel, err = h.createAtom("CLIPBOARD"); err != nil {
		return err
	}

	return nil
}

func (h *x11Clipboard) getClipboardOwner(selection xproto.Atom) (xproto.Window, bool) {
	reply, err := xproto.GetSelectionOwner(h.getConnection(), selection).Reply()
	if err != nil {
		h.moduleLogger.Warn("Failed read x11 clipboard owner", zap.Error(err))
		return 0, false
	}

	if reply.Owner == xproto.AtomNone {
		return 0, false
	}

	return reply.Owner, true
}

func (h *x11Clipboard) writeToClipboard(selection xproto.Atom, data Mimes) {
	h.ownClipboardData[selection] = data
	xproto.SetSelectionOwner(h.getConnection(), h.getWindow(), selection, xproto.TimeCurrentTime)
}

func (h *x11Clipboard) convertSelection(selection xproto.Atom) bool {
	owner, ok := h.getClipboardOwner(selection)
	if !ok {
		return false
	}

	if owner == h.getWindow() {
		return false
	}

	xproto.ConvertSelection(
		h.getConnection(), h.getWindow(), selection, h.atomUTF8String, selection, xproto.TimeCurrentTime)

	return true
}

func (h *x11Clipboard) onStart() {
	xevent.HookFun(h.eventHandlerHook).Connect(h.xConnection)
}

func (h *x11Clipboard) eventHandlerHook(xu *xgbutil.XUtil, eventRaw interface{}) bool {
	switch event := eventRaw.(type) {
	case xproto.SelectionNotifyEvent:
		if event.Requestor == h.getWindow() {
			xu.TimeSet(event.Time)
			h.onSelectionNotify(event)
			return false
		}
	case xproto.SelectionRequestEvent:
		if event.Owner == h.getWindow() {
			xu.TimeSet(event.Time)
			h.onSelectionRequest(event)
			return false
		}
	case xproto.SelectionClearEvent:
		if event.Owner == h.getWindow() {
			xu.TimeSet(event.Time)
			h.onSelectionClear(event)
			return false
		}
	}

	return true
}

func (h *x11Clipboard) onSelectionNotify(event xproto.SelectionNotifyEvent) {
	// var ptyp xproto.Atom
	b := make([]byte, 0, 1024)

	bytesAfter := uint32(1)
	bufsz := uint32(0)
	for bytesAfter > 0 {
		// last two args are offset and amount to transfer, in 32bit "long" sizes
		prop, err := xproto.GetProperty(
			h.getConnection(), true, h.getWindow(), event.Property, xproto.AtomAny, bufsz/4, MB/4).Reply()
		if err != nil {
			h.moduleLogger.Error("Failed get property for read clipboard", zap.Error(err))
			return
		}
		bytesAfter = prop.BytesAfter
		sz := len(prop.Value)
		if sz > 0 {
			b = append(b, prop.Value...)
			bufsz += uint32(sz)
		}
		// ptyp = prop.Type
	}

	var result Mimes
	isMulti, mediaType, boundary, body := IsMultipart(b)
	if isMulti {
		result = FromMultipart(body, boundary)
	} else {
		if mediaType != "" {
			result = NewMime(mediaType, b)
		} else {
			// unknown media type, try use TextPlain
			result = NewMime(TextPlain, b)
		}
	}

	fmt.Println(result.Text(TextPlain))
}

func (h *x11Clipboard) onSelectionRequest(event xproto.SelectionRequestEvent) {
	reply := xproto.SelectionNotifyEvent{
		Time:      event.Time,
		Requestor: event.Requestor,
		Selection: event.Selection,
		Target:    event.Target,
		Property:  xproto.AtomNone,
	}

	mask := xproto.EventMaskNoEvent
	if clipboardData, ok := h.ownClipboardData[event.Selection]; ok {
		reply.Property = event.Property
		if reply.Property == xproto.AtomNone {
			reply.Property = reply.Target
		}
		switch reply.Target {
		case h.atomTargets:
			mask = xproto.EventMaskPropertyChange
			targs := make([]byte, 4*3)
			bi := 0
			xgb.Put32(targs[bi:], uint32(h.atomUTF8String))
			bi += 4
			xgb.Put32(targs[bi:], uint32(h.atomTimestamp))
			bi += 4
			xgb.Put32(targs[bi:], uint32(h.atomTargets))
			xproto.ChangeProperty(h.getConnection(), xproto.PropModeReplace, reply.Requestor,
				reply.Property, xproto.AtomAtom, 32, 3, targs)
		case h.atomTimestamp:
			mask = xproto.EventMaskPropertyChange
			targs := make([]byte, 4*1)
			xgb.Put32(targs, uint32(xproto.TimeCurrentTime))
			xproto.ChangeProperty(h.getConnection(), xproto.PropModeReplace, reply.Requestor,
				reply.Property, xproto.AtomInteger, 32, 1, targs)
		case h.atomUTF8String:
			mask = xproto.EventMaskPropertyChange
			if len(clipboardData) > 1 {
				mpd := clipboardData.ToMultipart()
				xproto.ChangeProperty(h.getConnection(), xproto.PropModeReplace, reply.Requestor,
					reply.Property, reply.Target, 8, uint32(len(mpd)), mpd)
			} else {
				d := clipboardData[0]
				xproto.ChangeProperty(h.getConnection(), xproto.PropModeReplace, reply.Requestor,
					reply.Property, reply.Target, 8, uint32(len(d.Data)), d.Data)
			}
		}
	}

	xproto.SendEvent(h.getConnection(), false, reply.Requestor, uint32(mask), string(reply.Bytes()))
}

func (h *x11Clipboard) onSelectionClear(event xproto.SelectionClearEvent) {
	fmt.Println("onSelectionClear")
}

func (h *x11Clipboard) onStop() {
	// ci.lastWrite = nil
	// xproto.SetSelectionOwner(theApp.xc, xproto.AtomNone, theApp.atomClipboardSel, xproto.TimeCurrentTime)
	// keybind.Detach(h.xConnection, h.getWindow())
}
