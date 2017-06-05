package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cobra-viper-test",
	Short: "Cobra Viper Test",
	Long: `Long description of cobra/viper test`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
