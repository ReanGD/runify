package x11

import (
	"fmt"
	"os"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	zapOnSelectionChange = zap.String("Method", "x11Clipboard::onSelectionChange")
	zapOnSelectionNotify = zap.String("Method", "x11Clipboard::onSelectionNotify")
	zapOnPropertyNotify  = zap.String("Method", "x11Clipboard::onPropertyNotify")
)

type readData struct {
	owner     xproto.Window
	target    xproto.Atom
	timestamp xproto.Timestamp
	finish    bool
	data      []byte
}

type writeData struct {
	data Mimes
}

type x11Clipboard struct {
	atoms             *atomStorage
	connection        *xgb.Conn
	window            xproto.Window
	atomIncr          xproto.Atom
	atomTargets       xproto.Atom
	atomTargetsProp   xproto.Atom
	atomTimestamp     xproto.Atom
	atomImagePng      xproto.Atom
	atomImageBmp      xproto.Atom
	atomTextPlain     xproto.Atom
	atomUTF8String    xproto.Atom
	atomPrimarySel    xproto.Atom
	atomPrimaryProp   xproto.Atom
	atomClipboardSel  xproto.Atom
	atomClipboardProp xproto.Atom
	lastRead          map[xproto.Atom]*readData
	lastWrite         map[xproto.Atom]writeData
	incrementalMode   bool

	moduleLogger *zap.Logger
}

func newX11Clipboard() *x11Clipboard {
	return &x11Clipboard{
		atoms:             nil,
		connection:        nil,
		window:            0,
		atomIncr:          0,
		atomTargets:       0,
		atomTargetsProp:   0,
		atomTimestamp:     0,
		atomImagePng:      0,
		atomImageBmp:      0,
		atomTextPlain:     0,
		atomUTF8String:    0,
		atomPrimarySel:    0,
		atomPrimaryProp:   0,
		atomClipboardSel:  0,
		atomClipboardProp: 0,
		lastRead:          make(map[xproto.Atom]*readData),
		lastWrite:         make(map[xproto.Atom]writeData),
		incrementalMode:   false,
		moduleLogger:      nil,
	}
}

func (h *x11Clipboard) onInit(atoms *atomStorage, connection *xgb.Conn, window xproto.Window, moduleLogger *zap.Logger) error {
	h.atoms = atoms
	h.connection = connection
	h.window = window
	h.moduleLogger = moduleLogger

	h.atomIncr = atoms.getByNameUnchecked(atomNameIncr)
	h.atomTargets = atoms.getByNameUnchecked(atomNameTargets)
	h.atomTargetsProp = atoms.getByNameUnchecked(atomNameTargetsProp)
	h.atomTimestamp = atoms.getByNameUnchecked(atomNameTimestamp)
	h.atomImagePng = atoms.getByNameUnchecked(atomNameImagePng)
	h.atomImageBmp = atoms.getByNameUnchecked(atomNameImageBmp)
	h.atomTextPlain = atoms.getByNameUnchecked(atomNameTextPlain)
	h.atomUTF8String = atoms.getByNameUnchecked(atomNameUTF8String)
	h.atomPrimarySel = atoms.getByNameUnchecked(atomNamePrimarySel)
	h.atomPrimaryProp = atoms.getByNameUnchecked(atomNamePrimaryProp)
	h.atomClipboardSel = atoms.getByNameUnchecked(atomNameClipboardSel)
	h.atomClipboardProp = atoms.getByNameUnchecked(atomNameClipboardProp)

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

func (h *x11Clipboard) readFinish(selection xproto.Atom) {
	readData := h.lastRead[selection]
	readData.finish = true
	if readData.target != h.atomImagePng && readData.target != h.atomImageBmp {
		fmt.Println("read clipboard text:", string(readData.data))
	} else {
		fmt.Println("read clipboard image:", len(readData.data))
		fileName := "~/tmp/clipboard.png"
		if readData.target == h.atomImageBmp {
			fileName = "~/tmp/clipboard.bmp"
		}
		openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			h.moduleLogger.Warn("Failed open file", zapOnSelectionNotify, zap.Error(err))
			return
		}

		if _, err := openFile.Write(readData.data); err != nil {
			openFile.Close()
			h.moduleLogger.Warn("Failed write file", zapOnSelectionNotify, zap.Error(err))
			return
		}
		openFile.Close()
	}
}

func (h *x11Clipboard) readProperty(property xproto.Atom, fields ...zap.Field) ([]byte, readPropertyResult) {
	deleteProperty := true
	result := make([]byte, 0, 1024)
	chunkSize := uint32(1024 * 1024) // 1 MB
	for {
		// convert size of bytes to size of longs
		longOffset := uint32(len(result)) / 4
		longLength := chunkSize / 4

		reply, err := xproto.GetProperty(
			h.connection, deleteProperty, h.window, property, xproto.AtomAny, longOffset, longLength).Reply()
		if err != nil {
			h.moduleLogger.Warn("Failed call x11 get property",
				append(fields, h.atoms.getZapFieldPrefix("Property", property), zap.Error(err))...)
			return []byte{}, rpFailed
		}

		if reply == nil {
			h.moduleLogger.Warn("Failed call x11 get property",
				append(fields, h.atoms.getZapFieldPrefix("Property", property))...)
			return []byte{}, rpFailed
		}

		if reply.Type == h.atomIncr {
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

func (h *x11Clipboard) getClipboardOwner(selection xproto.Atom, fields ...zap.Field) (xproto.Window, bool) {
	reply, err := xproto.GetSelectionOwner(h.connection, selection).Reply()
	if err != nil {
		h.moduleLogger.Warn("Failed call x11 get clipboard owner",
			append(fields, h.atoms.getZapFieldPrefix("Selection", selection), zap.Error(err))...)
		return 0, false
	}

	if reply.Owner == xproto.AtomNone {
		return 0, false
	}

	return reply.Owner, true
}

func (h *x11Clipboard) convertSelection(
	selection xproto.Atom, target xproto.Atom, property xproto.Atom, fields ...zap.Field) bool {

	err := xproto.ConvertSelectionChecked(
		h.connection, h.window, selection, target, property, xproto.TimeCurrentTime).Check()

	if err != nil {
		h.moduleLogger.Warn("Failed call x11 convert selection",
			append(fields,
				h.atoms.getZapFieldPrefix("Selection", selection),
				h.atoms.getZapFieldPrefix("Target", target),
				h.atoms.getZapFieldPrefix("Property", property),
				zap.Error(err),
			)...)
		return false
	}
	return true
}

// func (h *x11Clipboard) writeToClipboard(selection xproto.Atom, data Mimes) bool {
// 	h.lastWrite[selection] = writeData{data: data}
// 	h.lastRead[selection] = readData{data: data, owner: h.window}
// 	if err := xproto.SetSelectionOwnerChecked(h.connection, h.window, selection, xproto.TimeCurrentTime).Check(); err != nil {
// 		h.moduleLogger.Warn("Failed set x11 clipboard owner", zap.Error(err))
// 		return false
// 	}

// 	return true
// }

func (h *x11Clipboard) onSelectionChange(event xfixes.SelectionNotifyEvent) {
	selection := event.Selection
	if !h.atoms.checkSelection(selection, zapOnSelectionChange) {
		return
	}

	if event.Subtype != xfixes.SelectionEventSetSelectionOwner {
		h.moduleLogger.Info("Unknown subtype", zapOnSelectionChange, zap.Uint32("Subtype", uint32(event.Subtype)))
	}

	if event.Owner == h.window {
		return
	}

	if event.Owner == xproto.AtomNone {
		delete(h.lastRead, selection)
		return
	}

	if data, ok := h.lastRead[selection]; ok && data.owner == event.Owner && data.timestamp == event.SelectionTimestamp {
		// already readed
		return
	}

	h.lastRead[selection] = &readData{
		owner:     event.Owner,
		target:    xproto.AtomNone,
		timestamp: event.SelectionTimestamp,
		finish:    false,
		data:      nil,
	}

	if !h.convertSelection(selection, h.atomTargets, h.atomTargetsProp, zapOnSelectionChange) {
		delete(h.lastRead, selection)
		return
	}
}

func (h *x11Clipboard) onSelectionNotify(event xproto.SelectionNotifyEvent) {
	selection := event.Selection
	if !h.atoms.checkSelection(selection, zapOnSelectionNotify) {
		return
	}

	readData, ok := h.lastRead[selection]
	if !ok {
		h.moduleLogger.Info("Not found read context", zapOnSelectionNotify, h.atoms.getZapFieldPrefix("Selection", selection))
		return
	}

	if event.Target == h.atomTargets {
		data, result := h.readProperty(event.Property, zapOnSelectionNotify)
		if result != rpSuccess {
			return
		}

		targets := make(map[xproto.Atom]struct{})
		for offset := 0; offset < len(data); offset += 4 {
			targets[xproto.Atom(xgb.Get32(data[offset:]))] = struct{}{}
		}

		selectionProp := h.atomPrimaryProp
		if selection == h.atoms.atomClipboardSel {
			selectionProp = h.atomClipboardProp
		}

		if selection == h.atoms.atomClipboardSel {
			for _, target := range []xproto.Atom{h.atomImagePng, h.atomImageBmp} {
				if _, ok := targets[target]; ok {
					if !h.convertSelection(event.Selection, target, selectionProp) {
						delete(h.lastRead, selection)
					} else {
						readData.target = target
					}
					return
				}
			}
		}

		for _, target := range []xproto.Atom{h.atomUTF8String, h.atomTextPlain} {
			if _, ok := targets[target]; ok {
				if !h.convertSelection(event.Selection, target, selectionProp) {
					delete(h.lastRead, selection)
				} else {
					readData.target = target
				}
				return
			}
		}

		delete(h.lastRead, selection)
		h.moduleLogger.Debug("No known clipboard types found in targets",
			zapOnSelectionNotify, h.atoms.getZapFieldPrefix("Selection", selection))
		return
	}

	if event.Target == readData.target && event.Target != xproto.AtomNone {
		data, result := h.readProperty(event.Property, zapOnSelectionNotify)
		if result == rpFailed {
			delete(h.lastRead, selection)
		} else if result == rpSuccess {
			readData.data = data
			h.readFinish(selection)
		} else if result == rpIncremental && selection == h.atoms.atomClipboardSel {
			h.incrementalMode = true
		} else if result == rpIncremental && selection == h.atoms.atomPrimarySel {
			delete(h.lastRead, selection)
			h.moduleLogger.Info("Incremental data for primary clipboard not supported", zapOnSelectionNotify)
		}

		return
	}

	h.moduleLogger.Debug("Unknown event target",
		zapOnSelectionNotify,
		h.atoms.getZapFieldPrefix("Target", event.Target),
		h.atoms.getZapFieldPrefix("Selection", selection))
}

func (h *x11Clipboard) onPropertyNotify(event xproto.PropertyNotifyEvent) {
	if event.Window == h.window && event.State == xproto.PropertyNewValue && event.Atom == h.atomClipboardProp && h.incrementalMode {
		readData, ok := h.lastRead[h.atomClipboardSel]
		if !ok {
			h.incrementalMode = false
			h.moduleLogger.Warn("Not found read context for incremental mode", zapOnPropertyNotify)
			return
		}
		if readData.target == xproto.AtomNone {
			h.incrementalMode = false
			h.moduleLogger.Warn("Wrong read context state for incremental mode, target == AtomNone", zapOnPropertyNotify)
			return
		}
		if readData.finish {
			h.incrementalMode = false
			h.moduleLogger.Warn("Wrong read context state for incremental mode, reading finished", zapOnPropertyNotify)
			return
		}

		data, result := h.readProperty(event.Atom, zapOnPropertyNotify)
		if result == rpFailed || result == rpIncremental {
			h.incrementalMode = false
			delete(h.lastRead, h.atomClipboardSel)
			h.moduleLogger.Warn("Failed read clipboard data in incremental mode, wrong result", zapOnPropertyNotify)
			return
		}
		if len(data) == 0 {
			h.incrementalMode = false
			h.readFinish(h.atomClipboardSel)
		} else if readData.data == nil {
			readData.data = data
		} else {
			readData.data = append(readData.data, data...)
		}

		return
	}
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
