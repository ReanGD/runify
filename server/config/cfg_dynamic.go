package config

import (
	"github.com/ReanGD/runify/server/paths"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type SystemCfg struct {
	UIBinaryPath string
	RpcAddress   string
	MainLocale   string
	DopLocale    string
	Terminal     string
}

func (c *SystemCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("System", map[string]interface{}{
		"UIBinaryPath": "/opt/runify/runify-ui",
		"RpcAddress":   "/tmp/runify.socket",
		"MainLocale":   "en_US.UTF-8",
		"DopLocale":    "ru_RU.UTF-8",
		"Terminal":     "sh",
	})
}

type ShortcutsCfg struct {
	Root string
}

func (c *ShortcutsCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Shortcuts", map[string]interface{}{
		"Root": "Super+R",
	})
}

type LoggerCfg struct {
	// One of: stderr, stdout, <filename> (for example: runify.log)
	Output string

	// MaxSizeMb is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSizeMb int

	// MaxAgeDays is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAgeDays int

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAgeDays may still cause them to get
	// deleted.)
	MaxBackups int

	// text is one of: "debug", "info", "warn", "error", "dpanic", "panic", "fatal"
	Level           zapcore.Level
	LevelStacktrace zapcore.Level
	AddCallerInfo   bool

	// text is one of: "plain", "json"
	Format LoggerFormat

	// LocalTimeInBackupFilename determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTimeInBackupFilename bool

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool

	// Enable/Disable log rotation with internal application settings
	EnableRotate bool
}

func (c *LoggerCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Logger", map[string]interface{}{
		"Level":                     "info",
		"LevelStacktrace":           "error",
		"AddCallerInfo":             false,
		"Format":                    "plain",
		"Output":                    "$RUNIFY_CACHE_DIR/runify.log",
		"MaxSizeMb":                 100,
		"MaxAgeDays":                3,
		"MaxBackups":                0,
		"LocalTimeInBackupFilename": true,
		"Compress":                  true,
		"EnableRotate":              true,
	})
}

type CfgDynamic struct {
	System    *SystemCfg
	Shortcuts *ShortcutsCfg
	Logger    *LoggerCfg
}

func newCfgDynamic() *CfgDynamic {
	return &CfgDynamic{
		System:    &SystemCfg{},
		Shortcuts: &ShortcutsCfg{},
		Logger:    &LoggerCfg{},
	}
}

func (c *CfgDynamic) postProcess() {
	c.System.UIBinaryPath = paths.ExpandAll(c.System.UIBinaryPath)
	c.System.RpcAddress = paths.ExpandAll(c.System.RpcAddress)
	c.Logger.Output = paths.ExpandAll(c.Logger.Output)
}

func (c *CfgDynamic) setDefault(vp *viper.Viper) {
	c.System.setDefault(vp)
	c.Shortcuts.setDefault(vp)
	c.Logger.setDefault(vp)
}
