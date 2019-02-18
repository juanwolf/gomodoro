package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
var cancelContext, cancel = context.WithCancel(context.Background())

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

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signals:
			cancel()
		}
	}()

}

func startTimer(duration time.Duration, refreshRate time.Duration, message string, ctx context.Context) {
	timerChannel, doneChannel := timer.Start(duration, refreshRate, ctx)
	outputManager.Start(duration, refreshRate, message)
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
