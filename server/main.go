package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

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

var cfgFile string

func main() {
	runify := root.NewRunify()
	rootCmd := &cobra.Command{
		Use:   "runify",
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

			runify.Run(cfgFile, buildCfg)
		},
	}

	var versionCmd = &cobra.Command{
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

	execFilepath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	defCfgPath := path.Join(filepath.Dir(execFilepath), "config.json")
	flags.StringVarP(&cfgFile, "config", "c", defCfgPath, "Config file path")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Command line start error: %s\n", err)
		return
	}
}
