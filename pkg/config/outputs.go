package config

import (
	"github.com/juanwolf/gomodoro/pkg/outputs"
	"github.com/spf13/viper"
)

type OutputsConfig struct {
	Stdout *StdoutConfig `mapstructure:"stdout"`
	File   *FileConfig   `mapstructure:"file"`
	Slack  *SlackConfig  `mapstructure:"slack"`
}

func DefaultOutputsConfig() OutputsConfig {
	stdoutConfig := StdoutConfig{
		ShowPercent: false,
		Size:        80,
	}
	return OutputsConfig{
		Stdout: &stdoutConfig,
		File:   nil,
		Slack:  nil,
	}
}

type OutputConfig interface {
	IsActivated() bool
	Instantiate() *outputs.Output
}

func (o OutputsConfig) GetOutputsConfig() []OutputConfig {
	return []OutputConfig{o.Stdout, o.File, o.Slack}
}

func setOutputsDefaults() {
	setStdoutDefaults()
}

type StdoutConfig struct {
	ShowPercent bool `mapstructure:"show_percent"`
	Size        int  `mapstructure:"size"`
}

func (c *StdoutConfig) IsActivated() bool {
	return c != nil
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

type FileConfig struct {
	Path string `mapstructure:"path"`
}

func (c *FileConfig) IsActivated() bool {
	return c != nil
}

func (c FileConfig) Instantiate() *outputs.Output {
	file := outputs.File{
		Path: c.Path,
	}
	output := outputs.Output(&file)
	return &output
}

func setFileDefaults() {
	viper.SetDefault("outputs.file.path", "$HOME/.gomodoro")
	viper.SetDefault("outputs.file.activated", false)
}

type SlackConfig struct {
	Token        string `mapstructure:"token"`
	DoNotDisturb bool   `mapstructure:"do_not_disturb"`
	Emoji        string `mapstructure:"emoji"`
}

func (c *SlackConfig) IsActivated() bool {
	return c != nil
}

func (c SlackConfig) Instantiate() *outputs.Output {
	slack := outputs.NewSlack(c.Token, c.DoNotDisturb, c.Emoji)
	output := outputs.Output(slack)
	return &output
}

func setSlackDefaults() {
	viper.SetDefault("outputs.slack.activated", false)
	viper.SetDefault("outputs.slack.emoji", ":tomato:")
	viper.SetDefault("outputs.slack.do_not_disturb", false)
}
