package cmd

import (
	"github.com/spf13/cobra"
)

var breakCmd = &cobra.Command{
	Use:   "break",
	Short: "Start a new break",
	Long:  `start a new break`,
	Run: func(cmd *cobra.Command, args []string) {
		startTimer(configuration.BreakDuration, configuration.RefreshRate)
	},
}

func init() {
	rootCmd.AddCommand(breakCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
