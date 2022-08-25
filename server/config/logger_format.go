package config

import (
	"bytes"
	"errors"
	"fmt"
)

type LoggerFormat int8

const (
	PlainFormat LoggerFormat = iota
	JSONFormat
)

var errUnmarshalNilLoggerFormat = errors.New("can't unmarshal a nil *LoggerFormat")

// String returns a lower-case ASCII representation of the log level.
func (l LoggerFormat) String() string {
	switch l {
	case PlainFormat:
		return "plain"
	case JSONFormat:
		return "json"
	default:
		return fmt.Sprintf("LoggerFormat(%d)", l)
	}
}

func (l LoggerFormat) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *LoggerFormat) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLoggerFormat
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized logger format: %q", text)
	}
	return nil
}

func (l *LoggerFormat) unmarshalText(text []byte) bool {
	switch string(text) {
	case "plain", "PLAIN":
		*l = PlainFormat
	case "json", "JSON":
		*l = JSONFormat
	default:
		return false
	}
	return true
}
