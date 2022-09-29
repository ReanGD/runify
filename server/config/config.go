package config

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const ModuleName = "config"

var logModule = zap.String("module", ModuleName)

type logItem struct {
	lvl zapcore.Level
	msg string
}

type Config struct {
	vp           *viper.Viper
	cfg          *Configuration
	deferredLog  []logItem
	moduleLogger *zap.Logger
}

func New(buildCfg *BuildCfg) *Config {
	vp := viper.New()
	configuration := newConfiguration(buildCfg)
	configuration.setDefault(vp)
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vp.SetConfigType("json")
	vp.AutomaticEnv()
	vp.SetEnvPrefix("RUNIFY")

	return &Config{
		vp:           vp,
		cfg:          configuration,
		deferredLog:  []logItem{},
		moduleLogger: nil,
	}
}

func (c *Config) OnInit(cfgFilePath string) {
	c.vp.SetConfigFile(cfgFilePath)

	if err := c.vp.ReadInConfig(); err != nil {
		msg := "Couldn't read config from file (%s), so let's switch to the config from the env variables and default values"
		c.deferredLog = append(c.deferredLog, logItem{lvl: zap.WarnLevel, msg: fmt.Sprintf(msg, err)})
	}

	cfg := new(ConfigurationSaved)
	err := c.vp.Unmarshal(&cfg, zapLevelDecoder)
	if err != nil {
		msg := "Couldn't decode config from file (%s), so let's switch to the config from the env variables and default values"
		c.deferredLog = append(c.deferredLog, logItem{lvl: zap.WarnLevel, msg: fmt.Sprintf(msg, err)})
	} else {
		cfg.process()
		c.cfg.ConfigurationSaved = cfg
	}
}

func (c *Config) OnStart(ctx context.Context, wg *sync.WaitGroup, rootLogger *zap.Logger) <-chan error {
	wg.Add(1)
	ch := make(chan error)
	c.moduleLogger = rootLogger.With(logModule)
	for _, item := range c.deferredLog {
		switch item.lvl {
		case zap.ErrorLevel:
			c.moduleLogger.Error(item.msg)
		case zap.WarnLevel:
			c.moduleLogger.Warn(item.msg)
		case zap.InfoLevel:
			c.moduleLogger.Info(item.msg)
		case zap.DebugLevel:
			c.moduleLogger.Debug(item.msg)
		}
	}
	c.deferredLog = []logItem{}
	wg.Done()
	return ch
}

func (c *Config) AddVersionToLog(log *zap.Logger) {
	cfg := c.cfg.Build
	log.Info("Runify version",
		zap.String("Version", cfg.Version),
		zap.String("Commit", cfg.BuildCommit),
		zap.String("Build data", cfg.BuildDateTime),
	)
}

func (c *Config) Get() *Configuration {
	return c.cfg
}

func (c *Config) Save() error {
	if err := c.vp.WriteConfig(); err != nil {
		return fmt.Errorf("unable to write config file: %s", err)
	}

	return nil
}
