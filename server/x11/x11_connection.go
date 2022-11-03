package x11

import (
	"errors"

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

type x11Connection struct {
	connection    *xgb.Conn
	rootWindowID  xproto.Window
	dummyWindowID xproto.Window
	quit          bool // when true, the main event loop will stop gracefully
	moduleLogger  *zap.Logger
}

func newX11Connection() *x11Connection {
	return &x11Connection{
		connection:    nil,
		rootWindowID:  0,
		dummyWindowID: 0,
		quit:          false,
		moduleLogger:  nil,
	}
}

func (c *x11Connection) onInit(display string, moduleLogger *zap.Logger) error {
	c.moduleLogger = moduleLogger

	if !c.createConnection(display) {
		if c.connection != nil {
			c.connection.Close()
		}
		return errInitX11Connection
	}

	if !c.createWindow() {
		c.connection.Close()
		return errInitX11Connection
	}

	return nil
}

func (c *x11Connection) createConnection(display string) bool {
	var err error
	if c.connection, err = xgb.NewConnDisplay(display); err != nil {
		c.moduleLogger.Warn("Failed connect to x11 server", zap.Error(err))
		return false
	}

	if err = xinerama.Init(c.connection); err != nil {
		c.moduleLogger.Warn("Failed init xinerama extension", zap.Error(err))
		return false
	}

	// https://www.x.org/releases/X11R7.7/doc/fixesproto/fixesproto.txt
	if err := xfixes.Init(c.connection); err != nil {
		c.moduleLogger.Warn("Failed init xfixes extension", zap.Error(err))
		return false
	}

	xfixesVersion, err := xfixes.QueryVersion(c.connection, xfixesMajorVersion, xfixesMinorVersion).Reply()
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

	return true
}

func (c *x11Connection) createWindow() bool {
	setupInfo := xproto.Setup(c.connection)
	screen := setupInfo.DefaultScreen(c.connection)
	c.rootWindowID = screen.Root

	var err error
	// Create a general purpose graphics context
	var gc xproto.Gcontext
	if gc, err = xproto.NewGcontextId(c.connection); err != nil {
		c.moduleLogger.Warn("Failed create x11 graphics context id", zap.Error(err))
		return false
	}

	err = xproto.CreateGCChecked(
		c.connection, gc, xproto.Drawable(c.rootWindowID), xproto.GcForeground, []uint32{screen.WhitePixel}).Check()
	if err != nil {
		c.moduleLogger.Warn("Failed create x11 graphics context", zap.Error(err))
		return false
	}

	if c.dummyWindowID, err = xproto.NewWindowId(c.connection); err != nil {
		c.moduleLogger.Warn("Failed create x11 window id", zap.Error(err))
		return false
	}

	err = xproto.CreateWindowChecked(
		c.connection, screen.RootDepth, c.dummyWindowID, c.rootWindowID,
		-1000, -1000, 1, 1, 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwEventMask|xproto.CwOverrideRedirect,
		[]uint32{1, xproto.EventMaskPropertyChange}).Check()
	if err != nil {
		c.moduleLogger.Warn("Failed create x11 window", zap.Error(err))
		return false
	}

	err = xproto.MapWindowChecked(c.connection, c.dummyWindowID).Check()
	if err != nil {
		c.moduleLogger.Warn("Failed map x11 window", zap.Error(err))
		return false
	}

	return true
}

func (c *x11Connection) onStart(x11EventsCh chan interface{}) error {
	var ev xgb.Event
	var err xgb.Error
	var resultErr error
	for !c.quit {
		if ev, err = c.connection.WaitForEvent(); err != nil {
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
			x11EventsCh <- event
		case xproto.KeyReleaseEvent:
			x11EventsCh <- event
		case xfixes.SelectionNotifyEvent:
			x11EventsCh <- event
		case xproto.SelectionNotifyEvent:
			x11EventsCh <- event
		case xproto.PropertyNotifyEvent:
			x11EventsCh <- event
		case xproto.SelectionRequestEvent:
			x11EventsCh <- event
		case xproto.SelectionClearEvent:
			x11EventsCh <- event
		}
	}

	if c.connection != nil {
		c.connection.Close()
	}

	return resultErr
}
