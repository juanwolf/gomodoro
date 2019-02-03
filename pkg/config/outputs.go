package config

import (
	"github.com/juanwolf/tomato/pkg/outputs"
	"github.com/spf13/viper"
)

type OutputsConfig struct {
	Stdout StdoutConfig `mapstructure:"stdout"`
}

func DefaultOutputsConfig() OutputsConfig {
	stdoutConfig := StdoutConfig{
		ShowPercent:   false,
		Size:          80,
		Activated:     true,
		FinishMessage: "Well done! Have a break, have a kitkat.",
	}

	return OutputsConfig{
		Stdout: stdoutConfig,
	}
}

type OutputConfig interface {
	IsActivated() bool
	Instantiate() *outputs.Output
}

func (o OutputsConfig) GetOutputsConfig() []OutputConfig {
	return []OutputConfig{o.Stdout}
}

func setOutputsDefaults() {
	setStdoutDefaults()
}

type StdoutConfig struct {
	ShowPercent   bool   `mapstructure:"show_percent"`
	Size          int    `mapstructure:"size"`
	Activated     bool   `mapstructure:"activated"`
	FinishMessage string `mapstructure:"finish_message"`
}

func (c StdoutConfig) IsActivated() bool {
	return c.Activated
}

func (c StdoutConfig) Instantiate() *outputs.Output {
	stdout := outputs.Stdout{
		ShowPercent: c.ShowPercent,
		Size:        c.Size,
	}
	output := outputs.Output(&stdout)
	return &output

}

func setStdoutDefaults() {
	viper.SetDefault("outputs.stdout.show_percent", false)
	viper.SetDefault("outputs.stdout.size", 80)
	viper.SetDefault("outputs.stdout.activated", true)
	viper.SetDefault("outputs.stdout.finish_message", "Well done. Have a break and let's get more stuff done!")
}
