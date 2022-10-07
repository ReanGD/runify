package x11

import (
	"fmt"

	"github.com/ReanGD/runify/server/system/mime"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	zapWriteToClipboard   = zap.String("Method", "x11Clipboard::writeToClipboard")
	zapOnSelectionChange  = zap.String("Method", "x11Clipboard::onSelectionChange")
	zapOnSelectionNotify  = zap.String("Method", "x11Clipboard::onSelectionNotify")
	zapOnPropertyNotify   = zap.String("Method", "x11Clipboard::onPropertyNotify")
	zapOnSelectionRequest = zap.String("Method", "x11Clipboard::onSelectionRequest")
	zapOnSelectionClear   = zap.String("Method", "x11Clipboard::onSelectionClear")
)

type x11Clipboard struct {
	atoms             *atomStorage
	connection        *xgb.Conn
	window            xproto.Window
	atomIncr          xproto.Atom
	atomTargets       xproto.Atom
	atomTargetsProp   xproto.Atom
	atomTimestamp     xproto.Atom
	atomPrimarySel    xproto.Atom
	atomPrimaryProp   xproto.Atom
	atomClipboardSel  xproto.Atom
	atomClipboardProp xproto.Atom
	lastRead          map[xproto.Atom]*readData
	lastWrite         map[xproto.Atom]*writeData
	incrReadMode      bool

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
		atomPrimarySel:    0,
		atomPrimaryProp:   0,
		atomClipboardSel:  0,
		atomClipboardProp: 0,
		lastRead:          make(map[xproto.Atom]*readData),
		lastWrite:         make(map[xproto.Atom]*writeData),
		incrReadMode:      false,
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

func (h *x11Clipboard) onStop() {

}

func (h *x11Clipboard) readFinish(selection xproto.Atom) {
	readData := h.lastRead[selection]
	readData.finish = true
	data := readData.data
	if data.IsText() {
		fmt.Println("read clipboard text:", string(data.Data))
	} else {
		fmt.Println("read clipboard image:")

		var fileExt string
		switch data.Type {
		case mime.ImagePng:
			fileExt = "png"
		case mime.ImageBmp:
			fileExt = "bmp"
		case mime.ImageJpeg:
			fileExt = "jpeg"
		}

		if err := data.WriteToFile(fmt.Sprintf("~/tmp/clipboard.%s", fileExt)); err != nil {
			h.moduleLogger.Warn("Failed write to file", zapOnSelectionNotify, zap.Error(err))
		}
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

func (h *x11Clipboard) writeProperty(
	window xproto.Window, property xproto.Atom, target xproto.Atom, valueLen byte, data []byte, fields ...zap.Field) {

	format := valueLen * 8
	dataLen := uint32(len(data) / int(valueLen))
	err := xproto.ChangePropertyChecked(
		h.connection, xproto.PropModeReplace, window, property, target, format, dataLen, data).Check()
	if err != nil {
		h.moduleLogger.Warn("Failed call x11 change property",
			append(fields, h.atoms.getZapFieldPrefix("Property", property), zap.Error(err))...)
	}
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

func (h *x11Clipboard) setSelectionOwner(selection xproto.Atom, fields ...zap.Field) bool {
	err := xproto.SetSelectionOwnerChecked(h.connection, h.window, selection, xproto.TimeCurrentTime).Check()
	if err != nil {
		h.moduleLogger.Warn("Failed call x11 set selection owner",
			append(fields,
				h.atoms.getZapFieldPrefix("Selection", selection),
				zap.Error(err),
			)...)
		return false
	}

	return true
}

func (h *x11Clipboard) writeToClipboard(isPrimary bool, data *mime.Data) bool {
	selection := h.atoms.atomClipboardSel
	if isPrimary {
		selection = h.atoms.atomPrimarySel
	}
	targetAtoms := h.atoms.getTargetAtomsByMime(data.Type)
	if len(targetAtoms) == 0 {
		h.moduleLogger.Warn("Failed write to clipboard, mime type not supported",
			data.Type.ZapField(), zapWriteToClipboard)
		return false
	}

	h.lastWrite[selection] = newWriteData(data, targetAtoms)
	return h.setSelectionOwner(selection, zapWriteToClipboard)
}

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

	h.lastRead[selection] = newReadData(event.Owner, event.SelectionTimestamp)

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

		if res, ok := h.atoms.choiceTarget(selection, targets); ok {
			selectionProp := h.atomPrimaryProp
			if selection == h.atoms.atomClipboardSel {
				selectionProp = h.atomClipboardProp
			}

			if res.mType.IsImage() && selection == h.atoms.atomPrimarySel {
				// Primary selection is not support image
				delete(h.lastRead, selection)
			} else if !h.convertSelection(event.Selection, res.atom, selectionProp) {
				delete(h.lastRead, selection)
			} else {
				readData.data = mime.NewEmptyData(res.mType)
			}

			return
		}

		delete(h.lastRead, selection)
		h.moduleLogger.Debug("No known clipboard types found in targets",
			zapOnSelectionNotify, h.atoms.getZapFieldPrefix("Selection", selection))
		return
	}

	if h.atoms.checkSelectionNotifyTarget(event.Target, readData.data.Type) {
		data, result := h.readProperty(event.Property, zapOnSelectionNotify)
		if result == rpFailed {
			delete(h.lastRead, selection)
		} else if result == rpSuccess {
			readData.data.Append(data)
			h.readFinish(selection)
		} else if result == rpIncremental && selection == h.atoms.atomClipboardSel {
			h.incrReadMode = true
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
	if event.Window == h.window && event.State == xproto.PropertyNewValue && event.Atom == h.atomClipboardProp && h.incrReadMode {
		readData, ok := h.lastRead[h.atomClipboardSel]
		if !ok {
			h.incrReadMode = false
			h.moduleLogger.Warn("Not found read context for incremental mode", zapOnPropertyNotify)
			return
		}
		if readData.data.Type == mime.None {
			h.incrReadMode = false
			h.moduleLogger.Warn("Wrong read context state for incremental mode, mType == mime.None", zapOnPropertyNotify)
			return
		}
		if readData.finish {
			h.incrReadMode = false
			h.moduleLogger.Warn("Wrong read context state for incremental mode, reading finished", zapOnPropertyNotify)
			return
		}

		data, result := h.readProperty(event.Atom, zapOnPropertyNotify)
		if result == rpFailed || result == rpIncremental {
			h.incrReadMode = false
			delete(h.lastRead, h.atomClipboardSel)
			h.moduleLogger.Warn("Failed read clipboard data in incremental mode, wrong result", zapOnPropertyNotify)
			return
		}
		if len(data) == 0 {
			h.incrReadMode = false
			h.readFinish(h.atomClipboardSel)
		} else {
			readData.data.Append(data)
		}

		return
	}
}

func (h *x11Clipboard) onSelectionRequest(event xproto.SelectionRequestEvent) {
	selection := event.Selection
	property := event.Property
	if property == xproto.AtomNone {
		property = event.Target
	}

	success := true
	if !h.atoms.checkSelection(selection, zapOnSelectionRequest) {
		success = false
	} else if writeData, ok := h.lastWrite[selection]; !ok {
		success = false
		h.moduleLogger.Info("Not found write context", zapOnSelectionRequest, h.atoms.getZapFieldPrefix("Selection", selection))
	} else if event.Target == h.atomTargets {
		propData := make([]byte, 4*(len(writeData.targetAtoms)+2))

		offset := 0
		xgb.Put32(propData[offset:], uint32(h.atomTargets))
		offset += 4
		xgb.Put32(propData[offset:], uint32(h.atomTimestamp))
		offset += 4
		for _, atom := range writeData.targetAtoms {
			xgb.Put32(propData[offset:], uint32(atom))
			offset += 4
		}
		h.writeProperty(event.Requestor, property, xproto.AtomAtom, 4, propData, zapOnSelectionRequest)
	} else if event.Target == h.atomTimestamp {
		propData := make([]byte, 4*1)
		xgb.Put32(propData, uint32(xproto.TimeCurrentTime))
		h.writeProperty(event.Requestor, property, xproto.AtomInteger, 4, propData, zapOnSelectionRequest)
	} else if writeData.isTargetAtom(event.Target) {
		h.writeProperty(event.Requestor, property, event.Target, 1, writeData.data.Data, zapOnSelectionRequest)
	} else {
		success = false
		h.moduleLogger.Debug("Unknown event target",
			zapOnSelectionRequest,
			h.atoms.getZapFieldPrefix("Target", event.Target),
			h.atoms.getZapFieldPrefix("Selection", selection))
	}

	mask := xproto.EventMaskPropertyChange
	if !success {
		mask = xproto.EventMaskNoEvent
		property = xproto.AtomNone
	}

	reply := xproto.SelectionNotifyEvent{
		Time:      event.Time,
		Requestor: event.Requestor,
		Selection: event.Selection,
		Target:    event.Target,
		Property:  property,
	}

	xproto.SendEvent(h.connection, false, reply.Requestor, uint32(mask), string(reply.Bytes()))
}

func (h *x11Clipboard) onSelectionClear(event xproto.SelectionClearEvent) {
	selection := event.Selection
	if !h.atoms.checkSelection(selection, zapOnSelectionClear) {
		return
	}

	if _, ok := h.lastWrite[selection]; ok {
		delete(h.lastWrite, selection)
	}
}
