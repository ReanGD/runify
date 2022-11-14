package x11

import (
	"github.com/ReanGD/runify/server/global/mime"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

const (
	atomNameNone          atomName = "NONE"
	atomNameIncr          atomName = "INCR"
	atomNameTargets       atomName = "TARGETS"
	atomNameTargetsProp   atomName = "TARGETS_PROP"
	atomNameTimestamp     atomName = "TIMESTAMP"
	atomNameImagePng      atomName = "image/png"
	atomNameImageBmp      atomName = "image/bmp"
	atomNameImageJpeg     atomName = "image/jpeg"
	atomNameTextPlain     atomName = "text/plain"
	atomNameTextPlainUtf8 atomName = "text/plain;charset=utf-8"
	atomNameUTF8String    atomName = "UTF8_STRING"
	atomNamePrimarySel    atomName = "PRIMARY"
	atomNamePrimaryProp   atomName = "PRIMARY_PROP"
	atomNameClipboardSel  atomName = "CLIPBOARD"
	atomNameClipboardProp atomName = "CLIPBOARD_PROP"
)

var (
	zapInitAtomStorage = zap.String("Method", "x11.atomStorage::init")
)

var (
	mimeAtomsMap = []struct {
		mType mime.Type
		name  atomName
	}{
		{mime.ImagePng, atomNameImagePng},
		{mime.ImageBmp, atomNameImageBmp},
		{mime.ImageJpeg, atomNameImageJpeg},
		{mime.TextPlain, atomNameUTF8String},
		{mime.TextPlain, atomNameTextPlain},
		{mime.TextPlain, atomNameTextPlainUtf8},
	}
)

type mimeAtom struct {
	mType mime.Type
	atom  xproto.Atom
}

type atomStorage struct {
	atomsByName       map[atomName]xproto.Atom
	atomsByID         map[xproto.Atom]atomName
	conn              *connection
	moduleLogger      *zap.Logger
	mimeAtoms         []mimeAtom
	atomIncr          xproto.Atom
	atomTargets       xproto.Atom
	atomTargetsProp   xproto.Atom
	atomTimestamp     xproto.Atom
	atomPrimarySel    xproto.Atom
	atomPrimaryProp   xproto.Atom
	atomClipboardSel  xproto.Atom
	atomClipboardProp xproto.Atom
}

func newAtomStorage() *atomStorage {
	return &atomStorage{
		atomsByName:       make(map[atomName]xproto.Atom),
		atomsByID:         make(map[xproto.Atom]atomName),
		conn:              nil,
		moduleLogger:      nil,
		mimeAtoms:         make([]mimeAtom, 0, len(mimeAtomsMap)),
		atomIncr:          xproto.AtomNone,
		atomTargets:       xproto.AtomNone,
		atomTargetsProp:   xproto.AtomNone,
		atomTimestamp:     xproto.AtomNone,
		atomPrimarySel:    xproto.AtomNone,
		atomPrimaryProp:   xproto.AtomNone,
		atomClipboardSel:  xproto.AtomNone,
		atomClipboardProp: xproto.AtomNone,
	}
}

func (s *atomStorage) init(conn *connection, moduleLogger *zap.Logger) bool {
	s.conn = conn
	s.moduleLogger = moduleLogger

	s.atomsByName[atomNameNone] = xproto.AtomNone
	s.atomsByID[xproto.AtomNone] = atomNameNone

	for _, it := range mimeAtomsMap {
		if atom, ok := s.getByName(it.name, zapInitAtomStorage); ok {
			s.mimeAtoms = append(s.mimeAtoms, mimeAtom{it.mType, atom})
		} else {
			s.moduleLogger.Warn("Failed init X11 atom storage",
				zap.String("Reason", "not found atom by name"), it.name.ZapField())
			return false
		}
	}

	atomsMap := []struct {
		name atomName
		atom *xproto.Atom
	}{
		{atomNameIncr, &s.atomIncr},
		{atomNameTargets, &s.atomTargets},
		{atomNameTargetsProp, &s.atomTargetsProp},
		{atomNameTimestamp, &s.atomTimestamp},
		{atomNamePrimarySel, &s.atomPrimarySel},
		{atomNamePrimaryProp, &s.atomPrimaryProp},
		{atomNameClipboardSel, &s.atomClipboardSel},
		{atomNameClipboardProp, &s.atomClipboardProp},
	}

	var ok bool
	for _, it := range atomsMap {
		*it.atom, ok = s.getByName(it.name, zapInitAtomStorage)
		if !ok {
			s.moduleLogger.Warn("Failed init X11 atom storage",
				zap.String("Reason", "not found atom by name"), it.name.ZapField())
			return false
		}

	}

	return true
}

func (s *atomStorage) getByName(name atomName, fields ...zap.Field) (xproto.Atom, bool) {
	if id, ok := s.atomsByName[name]; ok {
		return id, true
	}

	if id, ok := s.conn.createAtom(name, fields...); ok {
		s.atomsByName[name] = id
		s.atomsByID[id] = name

		return id, true
	}

	return xproto.AtomNone, false
}

func (s *atomStorage) getByID(id xproto.Atom, fields ...zap.Field) (atomName, bool) {
	if name, ok := s.atomsByID[id]; ok {
		return name, true
	}

	if name, ok := s.conn.getAtom(id, fields...); ok {
		s.atomsByName[name] = id
		s.atomsByID[id] = name

		return name, true
	}

	return "", false
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

func (s *atomStorage) isValidSelection(selection xproto.Atom, fields ...zap.Field) bool {
	if selection != s.atomPrimarySel && selection != s.atomClipboardSel {
		s.moduleLogger.Warn("Unknown selection atom",
			append(fields,
				s.getZapField(selection),
			)...)
		return false
	}

	return true
}

func (s *atomStorage) getTargetAtomsByMime(mType mime.Type) []xproto.Atom {
	var res []xproto.Atom
	for _, atom := range s.mimeAtoms {
		if atom.mType == mType {
			res = append(res, atom.atom)
		}
	}

	return res
}

func (s *atomStorage) choiceTarget(selection xproto.Atom, targets map[xproto.Atom]struct{}) (mimeAtom, bool) {
	for _, it := range s.mimeAtoms {
		if _, ok := targets[it.atom]; ok {
			return it, true
		}
	}

	return mimeAtom{mime.None, xproto.AtomNone}, false
}

func (s *atomStorage) checkSelectionNotifyTarget(target xproto.Atom, mType mime.Type) bool {
	if mType == mime.None || target == xproto.AtomNone {
		return false
	}

	for _, it := range s.mimeAtoms {
		if it.mType == mType && it.atom == target {
			return true
		}
	}

	return false
}
