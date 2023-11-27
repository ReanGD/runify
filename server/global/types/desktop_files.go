package types

type DesktopFile interface {
	ID() string
	FilePath() string
	IconPath() string
	Name() string
	Exec() string
	InTerminal() bool
}

type DesktopFiles []DesktopFile
