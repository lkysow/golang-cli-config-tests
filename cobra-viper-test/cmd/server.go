package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "short server descrip",
	Long: `longer server descrip`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := viper.GetString("config")
		if config != "" {
			viper.SetConfigFile(config)
			if err := viper.ReadInConfig(); err != nil {
				return fmt.Errorf("invalid config: %s", err)
			}
		}
		fmt.Printf("config: %v\n", viper.GetViper().AllSettings())
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	// set up the config flag
	serverCmd.Flags().String("config", "", "config file")
	viper.BindPFlag("config", serverCmd.Flags().Lookup("config"))

	// set up a gh-username flag that can also come from environment variables
	serverCmd.Flags().String("gh-username", "", "gh username")
	viper.BindEnv("gh-username", "GH_USERNAME")
	viper.BindPFlag("gh-username", serverCmd.Flags().Lookup("gh-username"))
}
