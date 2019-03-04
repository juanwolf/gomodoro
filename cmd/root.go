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
	"strconv"
	"time"
)

var rootCmd = &cobra.Command{
	Use:     "gomodoro",
	Short:   "Gomodoro is an integrated pomodoro timer",
	Long:    `An Integrated Pomodoro timer to save you from all distractions possible the time of a pomodoro`,
	Version: "v0.2.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configuration = &config.Config{}
var outputManager = outputs.NewOutputManager()
var cancelContext, cancel = context.WithCancel(context.Background())
var configFile string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.gomodoro.toml)")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signals:
			cancel()
		}
	}()
}

func initConfig() {
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

func startTimer(duration time.Duration, refreshRate time.Duration, lockFile string, message string, ctx context.Context) {
	timerChannel, doneChannel := timer.Start(duration, refreshRate, ctx)
	err := createLock(lockFile)
	if err != nil {
		fmt.Println("Can't create lock file. If no gomodoro is running, feel free to delete", lockFile)
		os.Exit(1)
	}
	outputManager.Start(duration, refreshRate, message)
	for {
		select {
		case <-doneChannel:
			outputManager.End()
			deleteLock(lockFile)
			return
		case timeElapsed := <-timerChannel:
			timeLeft := duration - timeElapsed
			outputManager.Refresh(timeLeft)
		}
	}
}

func createLock(lockFile string) error {
	file, err := os.OpenFile(lockFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	pid := os.Getpid()
	_, err = fmt.Fprintf(file, "%d", pid)
	return err
}

func deleteLock(lockFile string) error {
	_, err := os.Stat(lockFile)
	if err != nil {
		return err
	}
	err = os.Remove(lockFile)
	return err
}

// As the lock contains the PID, readLock return the PID of the main process
func readLock(lockFile string) (int, error) {
	file, err := os.Open(lockFile)
	defer file.Close()
	if err != nil {
		return 0, err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	pid := make([]byte, fileStat.Size())
	_, err = file.Read(pid)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(pid))
}
