package x11

import (
	"github.com/ReanGD/runify/server/global"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

type window struct {
	conn *connection
	id   xproto.Window
}

func newWindow(conn *connection, id xproto.Window) *window {
	return &window{
		conn: conn,
		id:   id,
	}
}

func (w *window) sendEvent(eventMask int, event string, fields ...zap.Field) bool {
	return w.conn.sendEvent(w.id, eventMask, event, fields...)
}

func (w *window) grabKey(modifiers uint16, keycode xproto.Keycode, fields ...zap.Field) global.Error {
	return w.conn.grabKey(w.id, modifiers, keycode, fields...)
}

func (w *window) ungrabKey(modifiers uint16, keycode xproto.Keycode, fields ...zap.Field) bool {
	return w.conn.ungrabKey(w.id, modifiers, keycode, fields...)
}

func (w *window) subscribeToSelectionChanges(selection xproto.Atom, fields ...zap.Field) bool {
	return w.conn.subscribeToSelectionChanges(w.id, selection, fields...)
}

func (w *window) readProperty(property xproto.Atom, fields ...zap.Field) ([]byte, readPropertyResult) {
	return w.conn.readProperty(w.id, property, fields...)
}

func (w *window) writeProperty(
	property xproto.Atom, target xproto.Atom, itemLen byte, data []byte, fields ...zap.Field,
) bool {
	return w.conn.writeProperty(w.id, property, target, itemLen, data, fields...)
}

func (w *window) setSelectionOwner(selection xproto.Atom, fields ...zap.Field) bool {
	return w.conn.setSelectionOwner(w.id, selection, fields...)
}

func (w *window) convertSelection(selection xproto.Atom, target xproto.Atom, property xproto.Atom, fields ...zap.Field) bool {
	return w.conn.convertSelection(w.id, selection, target, property, fields...)
}
