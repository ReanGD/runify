package x11

import "errors"

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
