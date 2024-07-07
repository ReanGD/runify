package portal

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/global/module"
	"github.com/ReanGD/runify/server/global/types"
	"github.com/godbus/dbus/v5"
	"go.uber.org/zap"
)

const (
	portalName = "org.freedesktop.portal.Desktop"
	portalPath = dbus.ObjectPath("/org/freedesktop/portal/desktop")

	portalRequestName           = "org.freedesktop.portal.Request"
	portalRequestMemberResponse = "Response"
	portalRequestSignalResponse = portalRequestName + "." + portalRequestMemberResponse
)

type dbusClient struct {
	conn              *dbus.Conn
	signalsCh         chan *dbus.Signal
	busObj            dbus.BusObject
	signalCallbacks   map[string]dbusSignalCallback
	responseCallbacks map[dbus.ObjectPath]dbusResponseCallback

	errorCtx     *module.ErrorCtx
	moduleLogger *zap.Logger
}

func newDBusClient(
	errorCtx *module.ErrorCtx,
	cfg *config.XDGDesktopPortalCfg,
	moduleLogger *zap.Logger,
) (*dbusClient, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		moduleLogger.Error("Failed to connect to session dbus", zap.Error(err))
		return nil, initErr
	}

	return &dbusClient{
		conn:              conn,
		signalsCh:         make(chan *dbus.Signal, cfg.SignalsChLen),
		busObj:            conn.Object(portalName, portalPath),
		signalCallbacks:   make(map[string]dbusSignalCallback),
		responseCallbacks: make(map[dbus.ObjectPath]dbusResponseCallback),
		errorCtx:          errorCtx,
		moduleLogger:      moduleLogger,
	}, nil
}

func (c *dbusClient) start() *types.HandledChannel {
	opts := []dbus.MatchOption{
		dbus.WithMatchInterface(portalRequestName),
		dbus.WithMatchMember(portalRequestMemberResponse),
	}

	if err := c.conn.AddMatchSignal(opts...); err != nil {
		c.moduleLogger.Error("Failed to add match dbus signal", zap.Error(err))
		c.errorCtx.SendError(startErr)
		return nil
	}

	c.conn.Signal(c.signalsCh)

	return types.NewHandledChannel(c.signalsCh, c.onEvent)
}

func (c *dbusClient) newToken() string {
	return fmt.Sprintf("runify_%d", rand.Int())
}

func (c *dbusClient) newObject(path dbus.ObjectPath) dbus.BusObject {
	return c.conn.Object(portalName, path)
}

func (c *dbusClient) signalSubscribe(name string, callback dbusSignalCallback) bool {
	if _, ok := c.signalCallbacks[name]; ok {
		return false
	}

	c.signalCallbacks[name] = callback
	return true
}

func (c *dbusClient) request(method string, token string, callback dbusResponseCallback, args ...any) error {
	busName := c.conn.Names()[0]
	busName = strings.ReplaceAll(busName, `:`, ``)
	busName = strings.ReplaceAll(busName, `.`, `_`)
	requestPath := dbus.ObjectPath(fmt.Sprintf("/org/freedesktop/portal/desktop/request/%s/%s", busName, token))

	c.responseCallbacks[requestPath] = callback
	return c.busObj.Call(method, 0, args...).Err
}

func (c *dbusClient) onEvent(event interface{}) (bool, error) {
	if signal, ok := event.(*dbus.Signal); ok {
		c.onSignal(signal)
	} else {
		c.moduleLogger.Warn("Failed cast event to dbus signal")
	}

	return false, nil
}

func (c *dbusClient) onSignal(signal *dbus.Signal) {
	name := signal.Name
	if name == portalRequestSignalResponse {
		if callback, ok := c.responseCallbacks[signal.Path]; ok {
			delete(c.responseCallbacks, signal.Path)
			if err := callback(signal.Body); err != nil {
				c.errorCtx.SendError(err)
			}
		} else {
			c.moduleLogger.Info("Unknown dbus response signal", zap.String("path", string(signal.Path)))
		}

		return
	}

	if callback, ok := c.signalCallbacks[name]; ok {
		callback(signal.Body)
		return
	}

	c.moduleLogger.Info("Unknown dbus signal", zap.String("signal", name))
}

func (c *dbusClient) close() {
	if c.conn == nil {
		return
	}

	if err := c.conn.Close(); err != nil {
		c.moduleLogger.Warn("Failed to close dbus connection", zap.Error(err))
	}

	c.conn = nil
}
