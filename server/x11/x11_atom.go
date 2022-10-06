package x11

import (
	"fmt"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

type atomName string

func (n atomName) ZapField() zap.Field {
	return zap.String("AtomName", string(n))
}

func (n atomName) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"AtomName", string(n))
}

const (
	atomNameNone          atomName = "NONE"
	atomNameIncr          atomName = "INCR"
	atomNameTargets       atomName = "TARGETS"
	atomNameTargetsProp   atomName = "TARGETS_PROP"
	atomNameTimestamp     atomName = "TIMESTAMP"
	atomNameImagePng      atomName = "image/png"
	atomNameImageBmp      atomName = "image/bmp"
	atomNameTextPlain     atomName = "text/plain"
	atomNameUTF8String    atomName = "UTF8_STRING"
	atomNamePrimarySel    atomName = "PRIMARY"
	atomNamePrimaryProp   atomName = "PRIMARY_PROP"
	atomNameClipboardSel  atomName = "CLIPBOARD"
	atomNameClipboardProp atomName = "CLIPBOARD_PROP"
)

var (
	allAtoms = []atomName{
		atomNameIncr,
		atomNameTargets,
		atomNameTargetsProp,
		atomNameTimestamp,
		atomNameImagePng,
		atomNameImageBmp,
		atomNameTextPlain,
		atomNameUTF8String,
		atomNamePrimarySel,
		atomNamePrimaryProp,
		atomNameClipboardSel,
		atomNameClipboardProp,
	}
)

type atomStorage struct {
	atomsByName      map[atomName]xproto.Atom
	atomsByID        map[xproto.Atom]atomName
	connection       *xgb.Conn
	moduleLogger     *zap.Logger
	atomPrimarySel   xproto.Atom
	atomClipboardSel xproto.Atom
}

func newAtomStorage(connection *xgb.Conn, moduleLogger *zap.Logger) (*atomStorage, bool) {
	res := &atomStorage{
		atomsByName:  make(map[atomName]xproto.Atom),
		atomsByID:    make(map[xproto.Atom]atomName),
		connection:   connection,
		moduleLogger: moduleLogger,
	}

	for _, name := range allAtoms {
		if _, ok := res.createAtom(name); !ok {
			return nil, false
		}
	}
	res.atomsByName[atomNameNone] = xproto.AtomNone
	res.atomsByID[xproto.AtomNone] = atomNameNone

	res.atomPrimarySel = res.getByNameUnchecked(atomNamePrimarySel)
	res.atomClipboardSel = res.getByNameUnchecked(atomNameClipboardSel)

	return res, true
}

func (s *atomStorage) createAtom(name atomName) (xproto.Atom, bool) {
	r, err := xproto.InternAtom(s.connection, false, uint16(len(name)), string(name)).Reply()
	if err != nil {
		s.moduleLogger.Error("Failed create x11 atom", name.ZapField(), zap.Error(err))
		return 0, false
	}
	if r == nil {
		s.moduleLogger.Error("Failed create x11 atom", name.ZapField())
		return 0, false
	}
	s.atomsByName[name] = r.Atom
	s.atomsByID[r.Atom] = name

	return r.Atom, true
}

func (s *atomStorage) getByName(name atomName) (xproto.Atom, bool) {
	if res, ok := s.atomsByName[name]; ok {
		return res, true
	}

	if res, ok := s.createAtom(name); ok {
		return res, true
	}

	return 0, false
}

func (s *atomStorage) getByNameUnchecked(name atomName) xproto.Atom {
	if res, ok := s.getByName(name); ok {
		return res
	}

	msg := fmt.Sprintf("Failed get x11 atom %s", name)
	panic(msg)
}

func (s *atomStorage) getByID(id xproto.Atom) (atomName, bool) {
	if res, ok := s.atomsByID[id]; ok {
		return res, true
	}

	reply, err := xproto.GetAtomName(s.connection, id).Reply()
	if err != nil {
		s.moduleLogger.Warn("Failed get x11 atom name", zap.Uint32("AtomID", uint32(id)), zap.Error(err))
		return "", false
	}

	if reply == nil {
		s.moduleLogger.Warn("Failed get x11 atom name", zap.Uint32("AtomID", uint32(id)))
		return "", false
	}

	return atomName(reply.Name), true
}

func (s *atomStorage) getZapField(id xproto.Atom) zap.Field {
	if name, ok := s.getByID(id); ok {
		return name.ZapField()
	}

	return zap.Uint32("AtomID", uint32(id))
}

func (s *atomStorage) getZapFieldPrefix(prefix string, id xproto.Atom) zap.Field {
	if name, ok := s.getByID(id); ok {
		return name.ZapFieldPrefix(prefix)
	}

	return zap.Uint32(prefix+"AtomID", uint32(id))
}

func (s *atomStorage) checkSelection(selection xproto.Atom, fields ...zap.Field) bool {
	if selection != s.atomPrimarySel && selection != s.atomClipboardSel {
		s.moduleLogger.Warn("Unknown selection atom", append(fields, s.getZapField(selection))...)
		return false
	}

	return true
}
