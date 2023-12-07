package de

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/ReanGD/runify/server/paths"
	"github.com/rkoesters/xdg/keyfile"
	"go.uber.org/zap"
)

var (
	// errMissingType means that the desktop entry is missing the
	// Type key, which is always required.
	errMissingType = errors.New("missing entry type")

	// errMissingName means that the desktop entry is missing the
	// Name key, which is required by the types Application, Link,
	// and Directory.
	errMissingName = errors.New("missing entry name")

	// errMissingURL means that the desktop entry is missing the URL
	// key, which is required by the type Link.
	errMissingURL = errors.New("missing entry url")
)

// Type is the type of desktop entry.
type Type uint8

// These are the possible desktop entry types.
const (
	None Type = iota // No type. This is bad.
	Application
	Link
	Directory
	Unknown // Any unknown type.
)

// ParseType converts the given string s into a Type.
func ParseType(s string) Type {
	switch s {
	case None.String():
		return None
	case Application.String():
		return Application
	case Link.String():
		return Link
	case Directory.String():
		return Directory
	default:
		return Unknown
	}
}

// String returns the Type as a string.
func (t Type) String() string {
	switch t {
	case None:
		return ""
	case Application:
		return "Application"
	case Link:
		return "Link"
	case Directory:
		return "Directory"
	default:
		return "Unknown"
	}
}

const (
	groupDesktopEntry        = "Desktop Entry"
	groupDesktopActionPrefix = "Desktop Action "

	keyType            = "Type"
	keyVersion         = "Version"
	keyName            = "Name"
	keyGenericName     = "GenericName"
	keyNoDisplay       = "NoDisplay"
	keyComment         = "Comment"
	keyIcon            = "Icon"
	keyHidden          = "Hidden"
	keyOnlyShowIn      = "OnlyShowIn"
	keyNotShowIn       = "NotShowIn"
	keyDBusActivatable = "DBusActivatable"
	keyTryExec         = "TryExec"
	keyExec            = "Exec"
	keyPath            = "Path"
	keyTerminal        = "Terminal"
	keyActions         = "Actions"
	keyMimeType        = "MimeType"
	keyCategories      = "Categories"
	keyImplements      = "Implements"
	keyKeywords        = "Keywords"
	keyStartupNotify   = "StartupNotify"
	keyStartupWMClass  = "StartupWMClass"
	keyURL             = "URL"
)

type desktopFile struct {
	// The unique id
	id string

	// The full path to the desktop entry file
	filePath string

	// The full path to the icon file
	iconPath string

	// The type of desktop entry. It can be: Application, Link, or Directory.
	entryType Type

	// Specific name of the application, for example "Mozilla"
	name string

	// Specific names for search, with \n as separator
	searchNames string

	// Icon to display in file manager, menus, etc.
	icon string

	// Program to execute
	exec string

	// Whether the program should be run in a terminal window
	terminal bool

	// If entry is Link type, the URL to access
	url string
}

func newDesktopFile(
	id string,
	filepath string,
	mainLocale keyfile.Locale,
	dopLocale keyfile.Locale,
	mimeStorage *mimeStorage,
	logger *zap.Logger,
) *desktopFile {
	fh, err := os.Open(filepath)
	if err != nil {
		logger.Info("Error open desktop entry file", zap.String("path", filepath), zap.Error(err))
		return nil
	}

	defer fh.Close()

	kf, err := keyfile.New(fh)
	if err != nil {
		logger.Info("Error parse desktop entry file struct", zap.String("path", filepath), zap.Error(err))
		return nil
	}

	res := &desktopFile{
		id:       id,
		filePath: filepath,
		iconPath: "",
	}

	ok, key, err := res.readFields(kf, mainLocale, dopLocale, mimeStorage)
	if err != nil {
		logger.Info("Error parse desktop entry file fields",
			zap.String("path", filepath), zap.String("field", key), zap.Error(err))
		return nil
	}
	if !ok {
		return nil
	}

	return res
}

func (f *desktopFile) readFields(
	kf *keyfile.KeyFile,
	mainLocale keyfile.Locale,
	dopLocale keyfile.Locale,
	mimeStorage *mimeStorage,
) (bool, string, error) {
	// see full list in github.com/rkoesters/xdg/desktop/entry.go
	if kf.KeyExists(groupDesktopEntry, keyNoDisplay) {
		noDisplay, err := kf.Bool(groupDesktopEntry, keyNoDisplay)
		if err != nil {
			return false, keyNoDisplay, err
		}
		if noDisplay {
			return false, "", nil
		}
	}

	if kf.KeyExists(groupDesktopEntry, keyHidden) {
		hidden, err := kf.Bool(groupDesktopEntry, keyHidden)
		if err != nil {
			return false, keyHidden, err
		}
		if hidden {
			return false, "", nil
		}
	}

	searchNames := map[string]struct{}{}

	f.entryType = ParseType(kf.Value(groupDesktopEntry, keyType))

	var err error
	f.name, err = kf.LocaleStringWithLocale(groupDesktopEntry, keyName, mainLocale)
	if err != nil {
		return false, keyName, err
	}

	searchName, err := kf.LocaleStringWithLocale(groupDesktopEntry, keyName, dopLocale)
	if err == nil {
		searchNames[searchName] = struct{}{}
	}

	if kf.KeyExists(groupDesktopEntry, keyIcon) {
		f.icon, err = kf.LocaleStringWithLocale(groupDesktopEntry, keyIcon, mainLocale)
		if err != nil {
			return false, keyIcon, err
		}
	}

	if kf.KeyExists(groupDesktopEntry, keyExec) {
		f.exec, err = kf.String(groupDesktopEntry, keyExec)
		if err != nil {
			return false, keyExec, err
		}
	}

	if kf.KeyExists(groupDesktopEntry, keyTerminal) {
		f.terminal, err = kf.Bool(groupDesktopEntry, keyTerminal)
		if err != nil {
			return false, keyTerminal, err
		}
	}

	if kf.KeyExists(groupDesktopEntry, keyMimeType) {
		mimeType, err := kf.StringList(groupDesktopEntry, keyMimeType)
		if err != nil {
			return false, keyMimeType, err
		}
		mimeStorage.addDesktopFile(mimeType, f)
	}

	if kf.KeyExists(groupDesktopEntry, keyKeywords) {
		keywords, err := kf.LocaleStringListWithLocale(groupDesktopEntry, keyKeywords, mainLocale)
		if err == nil {
			for _, keyword := range keywords {
				searchNames[keyword] = struct{}{}
			}
		}
		keywords, err = kf.LocaleStringListWithLocale(groupDesktopEntry, keyKeywords, dopLocale)
		if err == nil {
			for _, keyword := range keywords {
				searchNames[keyword] = struct{}{}
			}
		}
	}

	if f.entryType == Link {
		f.url, err = kf.String(groupDesktopEntry, keyURL)
		if err != nil {
			return false, keyURL, err
		}
	}

	// Validate the entry.
	if f.entryType == None {
		return false, keyType, errMissingType
	}
	if f.entryType > None && f.entryType < Unknown && f.name == "" {
		return false, keyType, errMissingName
	}
	if f.entryType == Link && f.url == "" {
		return false, keyType, errMissingURL
	}

	// Build search names
	delete(searchNames, f.name)
	f.searchNames = f.name
	for searchName := range searchNames {
		f.searchNames += "\n" + searchName
	}

	return true, "", nil
}

func (f *desktopFile) ID() string {
	return f.id
}

func (f *desktopFile) FilePath() string {
	return f.filePath
}

func (f *desktopFile) IconPath() string {
	return f.iconPath
}

func (f *desktopFile) Name() string {
	return f.name
}

func (f *desktopFile) SearchNames() string {
	return f.searchNames
}

func (f *desktopFile) fillParamCode(code rune, urls []string, files []string) []string {
	switch code {
	case 'u':
		if len(urls) > 0 {
			return []string{urls[0]}
		}
	case 'U':
		return urls
	case 'f':
		if len(files) > 0 {
			return []string{files[0]}
		}
	case 'F':
		return files
	}

	return []string{}
}

func (f *desktopFile) buildLaunchArgs(terminalPath string, urls []string, files []string) ([]string, error) {
	arg := []rune{}
	res := []string{}
	fieldCodeInd := -1
	inEscape := false
	inSingleQuote := false
	inDoubleQuote := false

	if f.terminal {
		if len(terminalPath) == 0 {
			return nil, fmt.Errorf("Terminal value is true, but terminal path is empty")
		}
		res = append(res, terminalPath, "-e")
	}

	for ind, c := range strings.Replace(f.exec, "\\\\", "\\", -1) {
		if inEscape {
			inEscape = false
			arg = append(arg, c)
			continue
		}

		switch c {
		case 'u', 'U', 'f', 'F':
			if fieldCodeInd == ind {
				if len(arg) > 0 {
					arg = arg[:len(arg)-1]
					res = append(res, f.fillParamCode(c, urls, files)...)
				}
				continue
			}
		case '"':
			if inDoubleQuote {
				inDoubleQuote = false
				res = append(res, string(arg))
				arg = arg[:0]
				continue
			}
			if !inSingleQuote {
				inDoubleQuote = true
				continue
			}

		case '\'':
			if inSingleQuote {
				inSingleQuote = false
				res = append(res, string(arg))
				arg = arg[:0]
				continue
			}
			if !inDoubleQuote {
				inSingleQuote = true
				continue
			}

		case '\\':
			if inDoubleQuote {
				inEscape = true
				continue
			}

		case '%':
			if !(inDoubleQuote || inSingleQuote) {
				fieldCodeInd = ind + 1
			}

		case ' ':
			if !(inDoubleQuote || inSingleQuote) {
				if len(arg) != 0 {
					res = append(res, string(arg))
					arg = arg[:0]
				}
				continue
			}
		}

		arg = append(arg, c)
	}

	if len(arg) != 0 {
		if !(inEscape || inDoubleQuote || inSingleQuote) {
			res = append(res, string(arg))
		} else {
			return nil, fmt.Errorf("Exec value contains an unbalanced number of quote characters: %s", f.exec)
		}
	}

	return res, nil
}

func (f *desktopFile) LaunchFull(terminalPath string, urls []string, files []string) error {
	args, err := f.buildLaunchArgs(terminalPath, urls, files)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		return errors.New("empty exec string")
	}

	name := args[0]
	var cmd *exec.Cmd
	if len(args) == 1 {
		cmd = exec.Command(name)
	} else {
		cmd = exec.Command(name, args[1:]...)
	}

	// If the parent process does not exit correctly, then all child processes will also be killed
	// This code cancel this behavior
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pgid: 0}
	cmd.Dir = paths.GetUserHome()

	if err = cmd.Start(); err != nil {
		return err
	}

	go cmd.Wait()

	return nil
}

func (f *desktopFile) Launch(terminalPath string) error {
	return f.LaunchFull(terminalPath, []string{}, []string{})
}

func (f *desktopFile) LaunchWithURLs(terminalPath string, urls ...string) error {
	return f.LaunchFull(terminalPath, urls, []string{})
}

func (f *desktopFile) LaunchWithFiles(terminalPath string, files ...string) error {
	return f.LaunchFull(terminalPath, []string{}, files)
}
