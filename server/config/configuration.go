package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type BuildCfg struct {
	Version       string
	BuildCommit   string
	BuildDateTime string
}

type ServerCfg struct {
	GrpcUnixAddr string
}

func (c ServerCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Server", map[string]interface{}{
		"GrpcUnixAddr": "/tmp/runify.socket",
	})
}

type ProviderCfg struct {
	ChannelLen          uint32
	SubModuleChannelLen uint32
	Hides               []string
}

func (c *ProviderCfg) setDefault(vp *viper.Viper) {
	vp.SetDefault("Provider", map[string]interface{}{
		"ChannelLen":          100,
		"SubModuleChannelLen": 100,
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
		"Output":                    "server.log",
		"MaxSizeMb":                 100,
		"MaxAgeDays":                3,
		"MaxBackups":                0,
		"LocalTimeInBackupFilename": true,
		"Compress":                  true,
		"EnableRotate":              true,
	})
}

type ConfigurationSaved struct {
	Server   ServerCfg
	Provider ProviderCfg
	Logger   LoggerCfg
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
	c.Server.setDefault(vp)
	c.Provider.setDefault(vp)
	c.Logger.setDefault(vp)
}