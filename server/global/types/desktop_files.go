package types

type DesktopFile interface {
	ID() string
	FilePath() string
	IconPath() string
	Name() string
	SearchNames() string

	LaunchFull(terminalPath string, urls []string, files []string) error
	Launch(terminalPath string) error
	LaunchWithURLs(terminalPath string, urls ...string) error
	LaunchWithFiles(terminalPath string, files ...string) error
}

type DesktopFiles []DesktopFile
