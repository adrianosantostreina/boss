package cmd

import (
	"os"

	"github.com/hashload/boss/cmd/config"
	"github.com/hashload/boss/core"
	"github.com/hashload/boss/env"
	"github.com/hashload/boss/msg"
	"github.com/hashload/boss/setup"
	"github.com/hashload/boss/utils"

	"github.com/spf13/cobra"
)

var versionPrint bool

var RootCmd = &cobra.Command{
	Use:   "boss",
	Short: "Dependency Manager for Delphi",
	Long:  "Dependency Manager for Delphi",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion(false)
	},
}

func Execute() {
	RootCmd.PersistentFlags().BoolVarP(&env.Global, "global", "g", false, "global environment")
	RootCmd.PersistentFlags().BoolVarP(&msg.DebugEnable, "debug", "d", false, "debug")
	RootCmd.Flags().BoolVarP(&versionPrint, "version", "v", false, "show cli version")

	msg.DebugEnable = utils.Contains(os.Args, "-d")

	setup.Initialize()

	config.InitializeConfig(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	core.RunGC(false)
}
