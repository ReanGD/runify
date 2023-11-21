package jdb

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/ReanGD/runify/server/paths"
	"go.uber.org/zap"
)

type JDB struct {
	fullPath string
	logger   *zap.Logger
}

func New(dbDir string, dbName string, logger *zap.Logger) (*JDB, error) {
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		logger.Error("Failed create directory for json database", zap.String("path", dbDir), zap.Error(err))
		return nil, errors.New("Failed open json database")
	}

	fullPath := filepath.Join(dbDir, dbName+".json")

	return &JDB{
		fullPath: fullPath,
		logger:   logger.With(zap.String("dbPath", fullPath)),
	}, nil
}

func (j *JDB) Read(dst interface{}) error {
	if ok, _ := paths.ExistsFile(j.fullPath); !ok {
		return nil
	}

	binData, err := os.ReadFile(j.fullPath)
	if err != nil {
		j.logger.Error("Failed read json database file", zap.Error(err))
		return errors.New("Failed read from json database")
	}

	err = json.Unmarshal(binData, dst)
	if err != nil {
		j.logger.Error("Failed unmarshal json database data", zap.Error(err))
		return errors.New("Failed read from json database")
	}

	return nil
}

func (j *JDB) Write(src interface{}) error {
	binData, err := json.Marshal(src)
	if err != nil {
		j.logger.Error("Failed marshal json database data", zap.Error(err))
		return errors.New("Failed write to json database")
	}

	if err = os.WriteFile(j.fullPath, binData, 0o644); err != nil {
		j.logger.Error("Failed write json database data", zap.Error(err))
		return errors.New("Failed write to json database")
	}

	return nil
}
