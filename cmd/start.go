// Copyright © 2019 Jean-Loup Adde
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/juanwolf/tomato/pkg/timer"
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
			case <-timerChannel:
				outputManager.Refresh()
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
