package cmd

import (
	"github.com/spf13/cobra"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start a new break",
	Long:  `start a new break`,
	Run: func(cmd *cobra.Command, args []string) {
		startTimer(configuration.BreakDuration, configuration.RefreshRate, configuration.LockFile, "", cancelContext)
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)
}
