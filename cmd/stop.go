package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"syscall"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running pomodoro",
	Long: `Stop the running pomodoro in case you started it in background.
It will actually just send a SIGTERM signal to the main process like you would do it if the process was in foreground.`,
	Run: func(cmd *cobra.Command, args []string) {
		pid, err := readLock(configuration.LockFile)
		if err != nil {
			fmt.Println("Can't read the lock file. Is there a pomodoro running?\nError:", err)
			os.Exit(1)
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println("Can't find running pomodoro! If no gomodoro is running, feel free to delete", configuration.LockFile, "\nError:", err)
			os.Exit(1)
		}
		process.Signal(syscall.SIGTERM)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
