package x11

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"go.uber.org/zap"
)

type atomName string

func (n atomName) ZapField() zap.Field {
	return zap.String("AtomName", string(n))
}

const (
	atomNameIncr         atomName = "INCR"
	atomNameTargets      atomName = "TARGETS"
	atomNameTimestamp    atomName = "TIMESTAMP"
	atomNameUTF8String   atomName = "UTF8_STRING"
	atomNamePrimarySel   atomName = "PRIMARY"
	atomNameClipboardSel atomName = "CLIPBOARD"
)

var (
	allAtoms = []atomName{
		atomNameIncr,
		atomNameTargets,
		atomNameTimestamp,
		atomNameUTF8String,
		atomNamePrimarySel,
		atomNameClipboardSel,
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

	panic("Atom not found")
}

func (s *atomStorage) getByID(id xproto.Atom) (atomName, bool) {
	if res, ok := s.atomsByID[id]; ok {
		return res, true
	}

	reply, err := xproto.GetAtomName(s.connection, id).Reply()
	if err != nil {
		s.moduleLogger.Error("Failed get x11 atom name", zap.Uint32("AtomID", uint32(id)), zap.Error(err))
		return "", false
	}

	if reply == nil {
		s.moduleLogger.Error("Failed get x11 atom name", zap.Uint32("AtomID", uint32(id)))
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

func (s *atomStorage) checkSelection(selection xproto.Atom, fields ...zap.Field) bool {
	if selection != s.atomPrimarySel && selection != s.atomClipboardSel {
		s.moduleLogger.Info("Unknown selection atom", append(fields, s.getZapField(selection))...)
		return false
	}

	return true
}
