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
	errInitX11Clipboard = errors.New("failed init X11 clipboard module")
)

type readPropertyResult uint8

const (
	rpFailed readPropertyResult = iota
	rpSuccess
	rpIncremental
)

type readData struct {
	finish    bool
	owner     xproto.Window
	timestamp xproto.Timestamp
	data      *mime.Data
}

func newReadData(owner xproto.Window, timestamp xproto.Timestamp) *readData {
	return &readData{
		finish:    false,
		owner:     owner,
		timestamp: timestamp,
		data:      nil,
	}
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
