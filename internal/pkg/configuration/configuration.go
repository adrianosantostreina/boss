package configuration

import (
	"github.com/spf13/cobra"
)

// Configuration is the struct of boss configuration
type Configuration struct {
	Global    bool
	DebugMode bool
	AppName   string
}

// InitializeConfiguration initialize the configuration of boss
func InitializeConfiguration(appName string) *Configuration {
	config := &Configuration{
		DebugMode: false,
		AppName:   appName,
	}
	return config
}

// BindFlags persist global command line flags
func (c *Configuration) BindFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&c.Global, "global", "g", false, "global environment")
	cmd.PersistentFlags().BoolVarP(&c.DebugMode, "debug", "d", false, "Debug mode")
}
