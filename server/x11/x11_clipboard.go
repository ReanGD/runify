package x11

import (
	"fmt"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	methodSelectionChange = zap.String("Method", "x11Clipboard::onSelectionChange")
)

type readData struct {
	owner     xproto.Window
	timestamp xproto.Timestamp
	data      Mimes
}

type writeData struct {
	data Mimes
}

type x11Clipboard struct {
	atoms            *atomStorage
	connection       *xgb.Conn
	window           xproto.Window
	atomIncr         xproto.Atom
	atomTargets      xproto.Atom
	atomTimestamp    xproto.Atom
	atomUTF8String   xproto.Atom
	atomPrimarySel   xproto.Atom
	atomClipboardSel xproto.Atom
	lastRead         map[xproto.Atom]readData
	lastWrite        map[xproto.Atom]writeData

	moduleLogger *zap.Logger
}

func newX11Clipboard() *x11Clipboard {
	return &x11Clipboard{
		atoms:            nil,
		connection:       nil,
		window:           0,
		atomIncr:         0,
		atomTargets:      0,
		atomTimestamp:    0,
		atomUTF8String:   0,
		atomPrimarySel:   0,
		atomClipboardSel: 0,
		lastRead:         make(map[xproto.Atom]readData),
		lastWrite:        make(map[xproto.Atom]writeData),
		moduleLogger:     nil,
	}
}

func (h *x11Clipboard) onInit(atoms *atomStorage, connection *xgb.Conn, window xproto.Window, moduleLogger *zap.Logger) error {
	h.atoms = atoms
	h.connection = connection
	h.window = window
	h.moduleLogger = moduleLogger

	h.atomIncr = atoms.getByNameUnchecked(atomNameIncr)
	h.atomTargets = atoms.getByNameUnchecked(atomNameTargets)
	h.atomTimestamp = atoms.getByNameUnchecked(atomNameTimestamp)
	h.atomUTF8String = atoms.getByNameUnchecked(atomNameUTF8String)
	h.atomPrimarySel = atoms.getByNameUnchecked(atomNamePrimarySel)
	h.atomClipboardSel = atoms.getByNameUnchecked(atomNameClipboardSel)

	// https://www.x.org/releases/X11R7.7/doc/fixesproto/fixesproto.txt
	if err := xfixes.Init(h.connection); err != nil {
		h.moduleLogger.Warn("Failed init xfixes extension", zap.Error(err))
		return errInitX11Clipboard
	}

	xfixesVersion, err := xfixes.QueryVersion(h.connection, xfixesMajorVersion, xfixesMinorVersion).Reply()
	if err != nil {
		h.moduleLogger.Warn("Failed get xfixes extension version", zap.Error(err))
		return errInitX11Clipboard
	}

	if (xfixesVersion.MajorVersion < xfixesMajorVersion) ||
		(xfixesVersion.MajorVersion == xfixesMajorVersion && xfixesVersion.MinorVersion < xfixesMinorVersion) {
		h.moduleLogger.Warn("Wrong xfixes extension version",
			zap.Uint32("expectedMajorVersion", xfixesMajorVersion),
			zap.Uint32("expectedMinorVersion", xfixesMinorVersion),
			zap.Uint32("actualMajorVersion", xfixesVersion.MajorVersion),
			zap.Uint32("actualMinorVersion", xfixesVersion.MinorVersion),
		)
		return errInitX11Clipboard
	}

	if err = xfixes.SelectSelectionInputChecked(
		h.connection, h.window, h.atomPrimarySel, xfixes.SelectionEventMaskSetSelectionOwner).Check(); err != nil {
		h.moduleLogger.Warn("Failed Subscribe to primary buffer changes", zap.Error(err))
		return errInitX11Clipboard
	}
	if err = xfixes.SelectSelectionInputChecked(
		h.connection, h.window, h.atomClipboardSel, xfixes.SelectionEventMaskSetSelectionOwner).Check(); err != nil {
		h.moduleLogger.Warn("Failed Subscribe to clipboard buffer changes", zap.Error(err))
		return errInitX11Clipboard
	}

	return nil
}

func (h *x11Clipboard) onStart() {

}

func (h *x11Clipboard) readProperty(property xproto.Atom, method zap.Field) ([]byte, bool) {
	// var ptyp xproto.Atom
	result := make([]byte, 0, 1024)

	bufsz := uint32(0)
	bytesAfter := uint32(1)
	chunkSize := uint32(1024 * 1024) // 1 MB
	for bytesAfter > 0 {
		delete := true

		// convert size of bytes to size of longs
		longOffset := bufsz / 4
		longLength := chunkSize / 4

		prop, err := xproto.GetProperty(
			h.connection, delete, h.window, property, xproto.AtomAny, longOffset, longLength).Reply()
		if err != nil {
			h.moduleLogger.Error("Failed get property for read clipboard", method, zap.Error(err))
			return []byte{}, false
		}

		bytesAfter = prop.BytesAfter
		sz := len(prop.Value)
		if sz > 0 {
			result = append(result, prop.Value...)
			bufsz += uint32(sz)
		}
		// ptyp = prop.Type
	}

	return result, true
}

func (h *x11Clipboard) getClipboardOwner(selection xproto.Atom) (xproto.Window, bool) {
	reply, err := xproto.GetSelectionOwner(h.connection, selection).Reply()
	if err != nil {
		h.moduleLogger.Warn("Failed read x11 clipboard owner", zap.Error(err))
		return 0, false
	}

	if reply.Owner == xproto.AtomNone {
		return 0, false
	}

	return reply.Owner, true
}

func (h *x11Clipboard) convertSelection(selection xproto.Atom, target xproto.Atom) bool {
	err := xproto.ConvertSelectionChecked(
		h.connection, h.window, selection, target, xproto.AtomNone, xproto.TimeCurrentTime).Check()

	if err != nil {
		h.moduleLogger.Warn("Failed convert x11 selection", zap.Error(err))
		return false
	}
	return true
}

func (h *x11Clipboard) writeToClipboard(selection xproto.Atom, data Mimes) bool {
	h.lastWrite[selection] = writeData{data: data}
	h.lastRead[selection] = readData{data: data, owner: h.window}
	if err := xproto.SetSelectionOwnerChecked(h.connection, h.window, selection, xproto.TimeCurrentTime).Check(); err != nil {
		h.moduleLogger.Warn("Failed set x11 clipboard owner", zap.Error(err))
		return false
	}

	return true
}

func (h *x11Clipboard) onSelectionChange(event xfixes.SelectionNotifyEvent) {
	selection := event.Selection
	if !h.atoms.checkSelection(selection, methodSelectionChange) {
		return
	}

	if event.Subtype != xfixes.SelectionEventSetSelectionOwner {
		h.moduleLogger.Info("Unknown subtype", methodSelectionChange, zap.Uint32("subtype", uint32(event.Subtype)))
	}

	if event.Owner == h.window {
		return
	}

	if event.Owner == xproto.AtomNone {
		delete(h.lastRead, selection)
		return
	}

	h.lastRead[selection] = readData{
		owner:     event.Owner,
		timestamp: event.Timestamp,
		data:      nil,
	}

	fmt.Println("onSelectionChange", event)
	if !h.convertSelection(selection, h.atomTargets) {
		return
	}
}

func (h *x11Clipboard) onSelectionNotify(event xproto.SelectionNotifyEvent) {
	method := zap.String("Method", "onSelectionNotify")
	if event.Selection != h.atomPrimarySel && event.Selection != h.atomClipboardSel {
		h.moduleLogger.Info("Unknown selection atom", method, zap.Uint32("selection", uint32(event.Selection)))
		return
	}

	if event.Target == h.atomTargets {
		data, ok := h.readProperty(event.Property, method)
		if !ok {
			return
		}

		for offset := 0; offset < len(data); offset += 4 {
			atom := xproto.Atom(xgb.Get32(data[offset:]))
			if atom == h.atomUTF8String {
				h.convertSelection(event.Selection, atom)
				return
			}
		}

		h.moduleLogger.Info("No known clipboard types found in targets", method, zap.Uint32("selection", uint32(event.Selection)))
		return
	}

	if event.Target == h.atomUTF8String {
		data, ok := h.readProperty(event.Property, method)
		if !ok {
			return
		}

		fmt.Println("read clipboard:", string(data))
		return
	}

	fmt.Println("onSelectionNotify", event)
	// if event.Target == h.atomTargets {
	// }

	// var result Mimes
	// isMulti, mediaType, boundary, body := IsMultipart(b)
	// if isMulti {
	// 	result = FromMultipart(body, boundary)
	// } else {
	// 	if mediaType != "" {
	// 		result = NewMime(mediaType, b)
	// 	} else {
	// 		// unknown media type, try use TextPlain
	// 		result = NewMime(TextPlain, b)
	// 	}
	// }

	// fmt.Println(result.Text(TextPlain))
}

func (h *x11Clipboard) onSelectionRequest(event xproto.SelectionRequestEvent) {
	// reply := xproto.SelectionNotifyEvent{
	// 	Time:      event.Time,
	// 	Requestor: event.Requestor,
	// 	Selection: event.Selection,
	// 	Target:    event.Target,
	// 	Property:  xproto.AtomNone,
	// }

	// mask := xproto.EventMaskNoEvent
	// if clipboardData, ok := h.ownClipboardData[event.Selection]; ok {
	// 	reply.Property = event.Property
	// 	if reply.Property == xproto.AtomNone {
	// 		reply.Property = reply.Target
	// 	}
	// 	switch reply.Target {
	// 	case h.atomTargets:
	// 		mask = xproto.EventMaskPropertyChange
	// 		targs := make([]byte, 4*3)
	// 		bi := 0
	// 		xgb.Put32(targs[bi:], uint32(h.atomUTF8String))
	// 		bi += 4
	// 		xgb.Put32(targs[bi:], uint32(h.atomTimestamp))
	// 		bi += 4
	// 		xgb.Put32(targs[bi:], uint32(h.atomTargets))
	// 		xproto.ChangeProperty(h.connection, xproto.PropModeReplace, reply.Requestor,
	// 			reply.Property, xproto.AtomAtom, 32, 3, targs)
	// 	case h.atomTimestamp:
	// 		mask = xproto.EventMaskPropertyChange
	// 		targs := make([]byte, 4*1)
	// 		xgb.Put32(targs, uint32(xproto.TimeCurrentTime))
	// 		xproto.ChangeProperty(h.connection, xproto.PropModeReplace, reply.Requestor,
	// 			reply.Property, xproto.AtomInteger, 32, 1, targs)
	// 	case h.atomUTF8String:
	// 		mask = xproto.EventMaskPropertyChange
	// 		if len(clipboardData) > 1 {
	// 			mpd := clipboardData.ToMultipart()
	// 			xproto.ChangeProperty(h.connection, xproto.PropModeReplace, reply.Requestor,
	// 				reply.Property, reply.Target, 8, uint32(len(mpd)), mpd)
	// 		} else {
	// 			d := clipboardData[0]
	// 			xproto.ChangeProperty(h.connection, xproto.PropModeReplace, reply.Requestor,
	// 				reply.Property, reply.Target, 8, uint32(len(d.Data)), d.Data)
	// 		}
	// 	}
	// }

	// xproto.SendEvent(h.connection, false, reply.Requestor, uint32(mask), string(reply.Bytes()))
}

func (h *x11Clipboard) onSelectionClear(event xproto.SelectionClearEvent) {
	// switch event.Selection {
	// case h.atomClipboardSel:
	// 	delete(h.ownClipboardData, h.atomClipboardSel)
	// case h.atomPrimarySel:
	// 	delete(h.ownClipboardData, h.atomPrimarySel)
	// default:
	// 	h.moduleLogger.Info("Unknown selection clear event", zap.Uint32("selection", uint32(event.Selection)))
	// }
}

func (h *x11Clipboard) onStop() {
	// ci.lastWrite = nil
	// xproto.SetSelectionOwner(theApp.xc, xproto.AtomNone, theApp.atomClipboardSel, xproto.TimeCurrentTime)
	// keybind.Detach(h.xConnection, h.getWindow())
}
