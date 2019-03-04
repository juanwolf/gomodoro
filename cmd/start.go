package cmd

import (
	"github.com/spf13/cobra"
)

var message string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new pomodoro",
	Long:  `start a new pomodoro for outputs specified in the config file`,
	Run: func(cmd *cobra.Command, args []string) {
		startTimer(configuration.PomodoroDuration, configuration.RefreshRate, configuration.LockFile, message, cancelContext)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&message, "message", "m", "", "Add a message to this pomodoro")
}
