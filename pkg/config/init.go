package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"time"
)

const (
	homeConfigFolder = "$HOME/.config/tomato"
	etcConfigFolder  = "/etc/tomato"
)

type Config struct {
	PomodoroDuration time.Duration `mapstructure:"pomodoro_duration"`
	RefreshRate      time.Duration `mapstructure:"refresh_rate"`
	Outputs          OutputsConfig `mapstructure:"outputs"`
}

func CopyDefaultConfig() error {
	destinationFile := "$HOME/.config/tomato/config.toml"
	input, err := ioutil.ReadFile("./config.toml")
	if err != nil {
		return err
	}

	if _, err := os.Stat(homeConfigFolder); os.IsNotExist(err) {
		os.Mkdir(homeConfigFolder, 0644)
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func setDefaults() {
	viper.SetDefault("pomodoro_duration", 25*time.Minute)
	viper.SetDefault("refresh_rate", 1*time.Second)
}

func ReadConfig(configPath string) (*Config, error) {
	var config Config
	viper.SetConfigType("toml")
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")         // name of config file (without extension)
		viper.AddConfigPath(etcConfigFolder)  // path to look for the config file in
		viper.AddConfigPath(homeConfigFolder) // call multiple times to add many search paths

		viper.AddConfigPath(".") // optionally look for config in the working directory
	}
	// Setting Default values in Viper
	setDefaults()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return nil, err
	}

	fmt.Println("Loading configuration from:", viper.ConfigFileUsed())

	err = viper.Unmarshal(&config)
	return &config, err
}
