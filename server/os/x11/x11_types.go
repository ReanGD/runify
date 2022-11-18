package x11

import (
	"errors"

	"github.com/ReanGD/runify/server/global/mime"
	"github.com/ReanGD/runify/server/global/shortcut"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

var (
	initErr = errors.New("failed to init x11 module")
)

type atomName string

func (n atomName) ZapField() zap.Field {
	return zap.String("AtomName", string(n))
}

func (n atomName) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"AtomName", string(n))
}

type readPropertyResult uint8

const (
	rpFailed readPropertyResult = iota
	rpSuccess
	rpIncremental
)

type readDataState uint8

const (
	rdsWaitType readDataState = iota
	rdsWaitData
	rdsReadIncr
	rdsFinished
)

type readData struct {
	state     readDataState
	owner     xproto.Window
	timestamp xproto.Timestamp
	data      *mime.Data
}

func newReadData(owner xproto.Window, timestamp xproto.Timestamp) *readData {
	return &readData{
		state:     rdsWaitType,
		owner:     owner,
		timestamp: timestamp,
		data:      nil,
	}
}

func (r *readData) setType(mType mime.Type) {
	r.state = rdsWaitData
	r.data = mime.NewEmptyData(mType)
}

func (r *readData) setIncrState() {
	r.state = rdsReadIncr
}
func (r *readData) finish() {
	r.state = rdsFinished
}

type writeData struct {
	data        *mime.Data
	targetAtoms []xproto.Atom
}

func newWriteData(data *mime.Data, targetAtoms []xproto.Atom) *writeData {
	return &writeData{
		data:        data,
		targetAtoms: targetAtoms,
	}
}

func (w *writeData) exists(atom xproto.Atom) bool {
	for _, targetAtom := range w.targetAtoms {
		if targetAtom == atom {
			return true
		}
	}
	return false
}

type x11Hotkey struct {
	mods    uint16
	keycode xproto.Keycode
}

type bindData struct {
	id         shortcut.HotkeyId
	x11Hotkeys []x11Hotkey
	hotkey     *shortcut.Hotkey
}

type subscribeToClipboardCmd struct {
	isPrimary bool
	ch        chan<- *mime.Data
	result    chan<- bool
}

func (c *subscribeToClipboardCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "subscribeToClipboard"),
		zap.Bool("IsPrimary", c.isPrimary),
		zap.String("Reason", reason),
		zap.String("Action", "subscription not activated"))
	c.result <- false
}

type writeToClipboardCmd struct {
	isPrimary bool
	data      *mime.Data
	result    chan<- bool
}

func (c *writeToClipboardCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "writeToClipboard"),
		zap.Bool("IsPrimary", c.isPrimary),
		zap.String("Reason", reason),
		zap.String("Action", "return empty result"))
	c.result <- false
}
