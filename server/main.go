package main

import (
	"fmt"
	"os"
	"path"

	"github.com/ReanGD/runify/server/config"
	"github.com/ReanGD/runify/server/root"
	"github.com/spf13/cobra"
)

var (
	version       string
	buildID       string
	buildUser     string
	buildCommit   string
	buildDateTime string
)

var (
	cfgSave bool
	cfgFile string
)

func main() {
	runify := root.NewRunify()
	rootCmd := &cobra.Command{
		Use:   "runify-server",
		Short: "Server part of runify",
		Long:  `Server part of runify.`,
		Run: func(cmd *cobra.Command, args []string) {
			buildCfg := &config.BuildCfg{
				Version:       version,
				BuildID:       buildID,
				BuildUser:     buildUser,
				BuildCommit:   buildCommit,
				BuildDateTime: buildDateTime,
			}

			runify.Run(cfgFile, cfgSave, buildCfg)
		},
	}

	versionCmd := &cobra.Command{
		Use:     "version",
		Short:   "Runify server version info",
		Long:    `Runify server version info`,
		Example: "runify-server version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version = %s\n", version)
			fmt.Printf("buildID = %s\n", buildID)
			fmt.Printf("buildUser = %s\n", buildUser)
			fmt.Printf("Commit = %s\n", buildCommit)
			fmt.Printf("Build data = %s\n", buildDateTime)
		},
	}

	rootCmd.AddCommand(versionCmd)

	flags := rootCmd.Flags()

	defCfgPath := os.ExpandEnv(path.Join("$XDG_CONFIG_HOME", "runify", "config.json"))
	flags.StringVarP(&cfgFile, "config", "c", defCfgPath, "Config file path")
	flags.BoolVarP(&cfgSave, "save", "s", false, "Save config to file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Command line start error: %s\n", err)
		return
	}
}
