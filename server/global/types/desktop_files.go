package types

type DesktopFile interface {
	ID() string
	FilePath() string
	IconPath() string
	Name() string
	SearchNames() string
	Exec() string
	InTerminal() bool
}

type DesktopFiles []DesktopFile
