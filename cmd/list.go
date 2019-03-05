package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the pomodoro",
	Long:  `List all the pomodoros inside the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		pomodoros, err := storeManager.GetPomodoros()
		if err != nil {
			panic(err)
		}
		for _, p := range pomodoros {
			fmt.Println(p.Id, "-", p.Created_at, "-", p.Message, "- stopped:", p.Stopped)
		}

		if len(pomodoros) == 0 {
			fmt.Println("No pomodoro in the database. Time to Work!")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
