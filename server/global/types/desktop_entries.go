package types

type DesktopEntry struct {
	// The path to the desktop entry file.
	filePath string

	// The path to the icon file.
	iconPath string

	// The real name of the desktop entry.
	name string
}

func NewDesktopEntry(filePath, iconPath, name string) *DesktopEntry {
	return &DesktopEntry{
		filePath: filePath,
		iconPath: iconPath,
		name:     name,
	}
}

func (e *DesktopEntry) FilePath() string {
	return e.filePath
}

func (e *DesktopEntry) IconPath() string {
	return e.iconPath
}

func (e *DesktopEntry) Name() string {
	return e.name
}

func (e *DesktopEntry) IsEqual(other *DesktopEntry) bool {
	return other != nil &&
		e.filePath == other.filePath &&
		e.iconPath == other.iconPath &&
		e.name == other.name
}

type DesktopEntries []*DesktopEntry

func (e DesktopEntries) Clone() DesktopEntries {
	result := make(DesktopEntries, len(e))
	for i, entry := range e {
		result[i] = entry
	}
	return result
}
