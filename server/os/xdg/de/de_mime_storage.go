package de

import "strings"

type mimeStorage struct {
	mimeTypes map[string][]*desktopFile
}

func newMimeStorage() *mimeStorage {
	return &mimeStorage{
		mimeTypes: map[string][]*desktopFile{},
	}
}

func (ms *mimeStorage) addDesktopFile(types []string, dfile *desktopFile) {
	for _, mimeType := range types {
		mimeType = strings.ToLower(strings.TrimSpace(mimeType))
		if mimeType != "" {
			ms.mimeTypes[mimeType] = append(ms.mimeTypes[mimeType], dfile)
		}
	}
}
