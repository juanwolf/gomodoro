package cmd

import (
	"fmt"
	"os"

	"github.com/juanwolf/gomodoro/pkg/config"
	"github.com/juanwolf/gomodoro/pkg/outputs"
	"github.com/juanwolf/gomodoro/pkg/timer"

	"github.com/spf13/cobra"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "tomato",
	Short: "Tomato is an integrated pomodoro timer",
	Long:  `An Integrated timer to be of all distractions possible the time of a pomodoro`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configuration = &config.Config{}
var outputManager = outputs.NewOutputManager()

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var configFile string
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.config/gomodoro/config.toml)")
	var err error
	configuration, err = config.ReadConfig(configFile)
	if err != nil {
		panic(err)
	}

	for _, o := range configuration.Outputs.GetOutputsConfig() {
		if o.IsActivated() {
			outputManager.Add(o.Instantiate())
		}
	}

}

func startTimer(duration time.Duration, refreshRate time.Duration) {
	timerChannel, doneChannel := timer.Start(duration, refreshRate)
	outputManager.Start(duration, refreshRate)
	for {
		select {
		case <-doneChannel:
			outputManager.End()
			return
		case timeElapsed := <-timerChannel:

			timeLeft := duration - timeElapsed
			outputManager.Refresh(timeLeft)
		}
	}
}
