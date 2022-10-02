package config

import (
	"github.com/ReanGD/runify/server/paths"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type BuildCfg struct {
	Version       string
	BuildCommit   string
	BuildDateTime string
}

type UICfg struct {
	BinaryPath string
}

func (c UICfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("UI", map[string]interface{}{
		"BinaryPath": "$RUNIFY_DATA_DIR/runify",
	})
}

type RpcCfg struct {
	ChannelLen uint32
	Address    string
}

func (c RpcCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Rpc", map[string]interface{}{
		"ChannelLen": 100,
		"Address":    "/tmp/runify.socket",
	})
}

type X11Cfg struct {
	ChannelLen         uint32
	HotkeysChannelLen  uint32
	X11EventChannelLen uint32
}

func (c X11Cfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("X11", map[string]interface{}{
		"ChannelLen":         100,
		"HotkeysChannelLen":  100,
		"X11EventChannelLen": 100,
	})
}

type ProviderCfg struct {
	ChannelLen          uint32
	SubModuleChannelLen uint32
	Terminal            string
	Hides               []string
}

func (c *ProviderCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Provider", map[string]interface{}{
		"ChannelLen":          100,
		"SubModuleChannelLen": 100,
		"Terminal":            "sh",
		"Hides":               []string{},
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

func (c LoggerCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Logger", map[string]interface{}{
		"Level":                     "info",
		"LevelStacktrace":           "error",
		"AddCallerInfo":             false,
		"Format":                    "plain",
		"Output":                    "runify.log",
		"MaxSizeMb":                 100,
		"MaxAgeDays":                3,
		"MaxBackups":                0,
		"LocalTimeInBackupFilename": true,
		"Compress":                  true,
		"EnableRotate":              true,
	})
}

type ConfigurationSaved struct {
	UI       UICfg
	Rpc      RpcCfg
	X11      X11Cfg
	Provider ProviderCfg
	Logger   LoggerCfg
}

func (c *ConfigurationSaved) process() {
	c.UI.BinaryPath = paths.ExpandAll(c.UI.BinaryPath)
	c.Rpc.Address = paths.ExpandAll(c.Rpc.Address)
}

type Configuration struct {
	Build *BuildCfg
	*ConfigurationSaved
}

func newConfiguration(buildCfg *BuildCfg) *Configuration {
	return &Configuration{
		Build:              buildCfg,
		ConfigurationSaved: new(ConfigurationSaved),
	}
}

func (c *Configuration) setDefault(vp *viper.Viper) {
	c.UI.setDefault(vp)
	c.Rpc.setDefault(vp)
	c.X11.setDefault(vp)
	c.Provider.setDefault(vp)
	c.Logger.setDefault(vp)
}
