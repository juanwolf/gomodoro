package cmd

import (
	"github.com/juanwolf/gomodoro/pkg/timer"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new pomodoro",
	Long:  `start a new pomodoro for outputs specified in the config file`,
	Run: func(cmd *cobra.Command, args []string) {
		timerChannel, doneChannel := timer.Start(configuration.PomodoroDuration, configuration.RefreshRate)
		outputManager.Start(configuration.PomodoroDuration, configuration.RefreshRate)
		for {
			select {
			case <-doneChannel:
				outputManager.End()
				return
			case timeElapsed := <-timerChannel:
				outputManager.Refresh(timeElapsed)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
