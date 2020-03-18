package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	LogLevel      string
	LogFormatter  string
	SpotifyConfig SpotifyConfig
}

// SpotifyConfig ...
type SpotifyConfig struct {
	ClientID    string
	SecretKey   string
	CallbackURL string
}

// ReadConfig ...
func ReadConfig() Config {
	var config Config

	// Set config name and type
	// Currently only YAML is supported
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Where we'll look for the file
	viper.AddConfigPath("/etc/syla/")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&config)

	return config
}
