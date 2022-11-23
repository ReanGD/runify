package config

type BuildCfg struct {
	Version       string
	BuildID       string
	BuildUser     string
	BuildCommit   string
	BuildDateTime string
}

type Configuration struct {
	Build *BuildCfg
	*CfgStatic
	*CfgDynamic
}

func newConfiguration(buildCfg *BuildCfg) *Configuration {
	return &Configuration{
		Build:      buildCfg,
		CfgStatic:  newCfgStatic(),
		CfgDynamic: newCfgDynamic(),
	}
}
