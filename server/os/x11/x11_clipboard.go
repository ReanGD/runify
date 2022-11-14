package x11

import (
	"github.com/ReanGD/runify/server/system/mime"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	zapInitClipboard      = zap.String("Method", "x11.clipboard::init")
	zapWriteToClipboard   = zap.String("Method", "x11.clipboard::writeToClipboard")
	zapOnSelectionChange  = zap.String("Method", "x11.clipboard::onSelectionChange")
	zapOnSelectionNotify  = zap.String("Method", "x11.clipboard::onSelectionNotify")
	zapOnPropertyNotify   = zap.String("Method", "x11.clipboard::onPropertyNotify")
	zapOnSelectionRequest = zap.String("Method", "x11.clipboard::onSelectionRequest")
	zapOnSelectionClear   = zap.String("Method", "x11.clipboard::onSelectionClear")
)

type clipboard struct {
	conn      *connection
	atoms     *atomStorage
	window    *window
	lastRead  map[xproto.Atom]*readData
	lastWrite map[xproto.Atom]*writeData

	moduleLogger *zap.Logger
}

func newClipboard() *clipboard {
	return &clipboard{
		conn:         nil,
		atoms:        nil,
		window:       nil,
		lastRead:     make(map[xproto.Atom]*readData),
		lastWrite:    make(map[xproto.Atom]*writeData),
		moduleLogger: nil,
	}
}

func (c *clipboard) init(conn *connection, atoms *atomStorage, window *window, moduleLogger *zap.Logger) bool {
	c.conn = conn
	c.atoms = atoms
	c.window = window
	c.moduleLogger = moduleLogger

	if !c.window.subscribeToSelectionChanges(c.atoms.atomPrimarySel, zapInitClipboard) {
		c.moduleLogger.Error("Failed init x11 cliboard", zap.String("Reason", "failed subscribe to primary buffer changes"))
		return false
	}

	if !c.window.subscribeToSelectionChanges(c.atoms.atomClipboardSel, zapInitClipboard) {
		c.moduleLogger.Error("Failed init x11 cliboard", zap.String("Reason", "failed subscribe to clipboard buffer changes"))
		return false
	}

	return true
}

func (c *clipboard) readFinish(selection xproto.Atom) {
	// readData := h.lastRead[selection]
	// data := readData.data
	// if data.IsText() {
	// 	fmt.Println("read clipboard text:", string(data.Data))
	// } else {
	// 	fmt.Println("read clipboard image:")

	// 	var fileExt string
	// 	switch data.Type {
	// 	case mime.ImagePng:
	// 		fileExt = "png"
	// 	case mime.ImageBmp:
	// 		fileExt = "bmp"
	// 	case mime.ImageJpeg:
	// 		fileExt = "jpeg"
	// 	}

	// 	if err := data.WriteToFile(fmt.Sprintf("~/tmp/clipboard.%s", fileExt)); err != nil {
	// 		h.moduleLogger.Warn("Failed write to file", zapOnSelectionNotify, zap.Error(err))
	// 	}
	// }
}

func (c *clipboard) writeToClipboard(isPrimary bool, data *mime.Data) bool {
	selection := c.atoms.atomClipboardSel
	if isPrimary {
		selection = c.atoms.atomPrimarySel
	}
	targetAtoms := c.atoms.getTargetAtomsByMime(data.Type)
	if len(targetAtoms) == 0 {
		c.moduleLogger.Warn("Failed write to clipboard, mime type not supported",
			data.Type.ZapField(), zapWriteToClipboard)
		return false
	}

	c.lastWrite[selection] = newWriteData(data, targetAtoms)
	return c.window.setSelectionOwner(selection, zapWriteToClipboard)
}

func (c *clipboard) onSelectionChange(event xfixes.SelectionNotifyEvent) {
	selection := event.Selection
	if !c.atoms.isValidSelection(selection, zapOnSelectionChange) {
		return
	}

	if event.Subtype != xfixes.SelectionEventSetSelectionOwner {
		c.moduleLogger.Info("Unknown subtype", zapOnSelectionChange, zap.Uint32("Subtype", uint32(event.Subtype)))
	}

	if event.Owner == c.window.id {
		return
	}

	if event.Owner == xproto.AtomNone {
		// selection was cleared
		delete(c.lastRead, selection)
		return
	}

	if data, ok := c.lastRead[selection]; ok && data.owner == event.Owner && data.timestamp == event.SelectionTimestamp {
		// already readed
		return
	}

	c.lastRead[selection] = newReadData(event.Owner, event.SelectionTimestamp)

	if !c.window.convertSelection(selection, c.atoms.atomTargets, c.atoms.atomTargetsProp, zapOnSelectionChange) {
		delete(c.lastRead, selection)
		return
	}
}

func (c *clipboard) onSelectionNotify(event xproto.SelectionNotifyEvent) {
	selection := event.Selection
	if !c.atoms.isValidSelection(selection, zapOnSelectionNotify) {
		return
	}

	readData, ok := c.lastRead[selection]
	if !ok {
		c.moduleLogger.Info("Not found read context", zapOnSelectionNotify, c.atoms.getZapFieldPrefix("Selection", selection))
		return
	}

	if readData.state == rdsWaitType && event.Target == c.atoms.atomTargets {
		data, result := c.window.readProperty(event.Property, zapOnSelectionNotify)
		if result != rpSuccess {
			return
		}

		targets := make(map[xproto.Atom]struct{})
		for offset := 0; offset < len(data); offset += 4 {
			targets[xproto.Atom(xgb.Get32(data[offset:]))] = struct{}{}
		}

		if res, ok := c.atoms.choiceTarget(selection, targets); ok {
			selectionProp := c.atoms.atomPrimaryProp
			if selection == c.atoms.atomClipboardSel {
				selectionProp = c.atoms.atomClipboardProp
			}

			if res.mType.IsImage() && selection == c.atoms.atomPrimarySel {
				// Primary selection is not support image
				delete(c.lastRead, selection)
			} else if !c.window.convertSelection(selection, res.atom, selectionProp) {
				delete(c.lastRead, selection)
			} else {
				readData.setType(res.mType)
			}

			return
		}

		delete(c.lastRead, selection)
		c.moduleLogger.Debug("No known clipboard types found in targets",
			zapOnSelectionNotify, c.atoms.getZapFieldPrefix("Selection", selection))
		return
	}

	if readData.state == rdsWaitData && c.atoms.checkSelectionNotifyTarget(event.Target, readData.data.Type) {
		data, result := c.window.readProperty(event.Property, zapOnSelectionNotify)
		if result == rpFailed {
			delete(c.lastRead, selection)
		} else if result == rpSuccess {
			readData.data.Append(data)
			readData.finish()
			c.readFinish(selection)
		} else if result == rpIncremental && selection == c.atoms.atomClipboardSel {
			readData.setIncrState()
		} else if result == rpIncremental && selection == c.atoms.atomPrimarySel {
			delete(c.lastRead, selection)
			c.moduleLogger.Info("Incremental data for primary clipboard not supported", zapOnSelectionNotify)
		}

		return
	}

	c.moduleLogger.Debug("Unknown event target",
		zapOnSelectionNotify,
		c.atoms.getZapFieldPrefix("Target", event.Target),
		c.atoms.getZapFieldPrefix("Selection", selection))
}

func (c *clipboard) onPropertyNotify(event xproto.PropertyNotifyEvent) {
	if event.Window == c.window.id && event.State == xproto.PropertyNewValue && event.Atom == c.atoms.atomClipboardProp {
		readData, ok := c.lastRead[c.atoms.atomClipboardSel]
		if !ok || readData.state != rdsReadIncr {
			return
		}

		data, result := c.window.readProperty(c.atoms.atomClipboardProp, zapOnPropertyNotify)
		if result == rpFailed || result == rpIncremental {
			delete(c.lastRead, c.atoms.atomClipboardSel)
			c.moduleLogger.Warn("Failed read clipboard data in incremental mode, wrong result", zapOnPropertyNotify)
		} else if len(data) == 0 {
			readData.finish()
			c.readFinish(c.atoms.atomClipboardSel)
		} else {
			readData.data.Append(data)
		}

		return
	}
}

func (c *clipboard) onSelectionRequest(event xproto.SelectionRequestEvent) {
	selection := event.Selection
	property := event.Property
	if property == xproto.AtomNone {
		property = event.Target
	}

	clientWin := c.conn.newWindow(event.Requestor)

	success := true
	if !c.atoms.isValidSelection(selection, zapOnSelectionRequest) {
		success = false
	} else if writeData, ok := c.lastWrite[selection]; !ok {
		success = false
		c.moduleLogger.Info("Not found write context", zapOnSelectionRequest, c.atoms.getZapFieldPrefix("Selection", selection))
	} else if event.Target == c.atoms.atomTargets {
		propData := make([]byte, 4*(len(writeData.targetAtoms)+2))

		offset := 0
		xgb.Put32(propData[offset:], uint32(c.atoms.atomTargets))
		offset += 4
		xgb.Put32(propData[offset:], uint32(c.atoms.atomTimestamp))
		offset += 4
		for _, atom := range writeData.targetAtoms {
			xgb.Put32(propData[offset:], uint32(atom))
			offset += 4
		}
		clientWin.writeProperty(property, xproto.AtomAtom, 4, propData, zapOnSelectionRequest)
	} else if event.Target == c.atoms.atomTimestamp {
		propData := make([]byte, 4*1)
		xgb.Put32(propData, uint32(xproto.TimeCurrentTime))
		clientWin.writeProperty(property, xproto.AtomInteger, 4, propData, zapOnSelectionRequest)
	} else if writeData.exists(event.Target) {
		clientWin.writeProperty(property, event.Target, 1, writeData.data.Data, zapOnSelectionRequest)
	} else {
		success = false
		c.moduleLogger.Debug("Unknown event target",
			zapOnSelectionRequest,
			c.atoms.getZapFieldPrefix("Target", event.Target),
			c.atoms.getZapFieldPrefix("Selection", selection))
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

	_ = clientWin.sendEvent(mask, string(reply.Bytes()), zapOnSelectionRequest)
}

func (c *clipboard) onSelectionClear(event xproto.SelectionClearEvent) {
	selection := event.Selection
	if !c.atoms.isValidSelection(selection, zapOnSelectionClear) {
		return
	}

	if _, ok := c.lastWrite[selection]; ok {
		delete(c.lastWrite, selection)
	}
}
