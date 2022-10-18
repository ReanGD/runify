package x11

import (
	"errors"

	"github.com/ReanGD/runify/server/system/mime"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

const (
	xfixesMajorVersion = 5
	xfixesMinorVersion = 0
)

var (
	errInitX11          = errors.New("failed init X11 server")
	errInitX11Keyboard  = errors.New("failed init X11 keyboard module")
	errInitX11Clipboard = errors.New("failed init X11 clipboard module")
)

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

func (w *writeData) isTargetAtom(atom xproto.Atom) bool {
	for _, targetAtom := range w.targetAtoms {
		if targetAtom == atom {
			return true
		}
	}
	return false
}

type writeToClipboardCmd struct {
	isPrimary bool
	data      *mime.Data
}

func (c *writeToClipboardCmd) onRequestDefault(logger *zap.Logger, reason string) {
	logger.Warn("Process message finished with error",
		zap.String("Request", "writeToClipboard"),
		zap.Bool("IsPrimary", c.isPrimary),
		zap.String("Reason", reason),
		zap.String("Action", "return empty result"))
}
