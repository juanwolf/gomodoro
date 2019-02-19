package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/user"
	"time"
)

const (
	homeConfigFolder = ".config/gomodoro"
)

type Config struct {
	PomodoroDuration time.Duration `mapstructure:"pomodoro_duration"`
	BreakDuration    time.Duration `mapstructure:"break_duration"`
	RefreshRate      time.Duration `mapstructure:"refresh_rate"`
	Outputs          OutputsConfig `mapstructure:"outputs"`
}

func CopyDefaultConfig() error {
	destinationFile := "$HOME/.gomodoro.toml"
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
	viper.SetDefault("break_duration", 5*time.Minute)
	viper.SetDefault("refresh_rate", 1*time.Second)
}

func ReadConfig(configPath string) (*Config, error) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	userConfig := fmt.Sprintf("%s/%s", usr.HomeDir, homeConfigFolder)
	var config Config
	viper.SetConfigType("toml")
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")   // name of config file (without extension)
		viper.AddConfigPath(userConfig) // call multiple times to add many search paths
		viper.SetConfigName(".gomodoro")
		viper.AddConfigPath(usr.HomeDir)
	}
	// Setting Default values in Viper
	setDefaults()
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return nil, err
	}

	fmt.Println("Loading configuration from:", viper.ConfigFileUsed())

	err = viper.Unmarshal(&config)
	return &config, err
}
