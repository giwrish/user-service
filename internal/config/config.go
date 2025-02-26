package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB     DatabaseConfig `mapstructure:"database"`
	Server ServerConfig   `mapstructure:"server"`
}

type ServerConfig struct {
	Port         int `mapstructure:"port"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	IdleTimeout  int `mapstructure:"idle_timeout"`
}

type DatabaseConfig struct {
	Url               string `mapstructure:"url"`
	MaxConnection     int32  `mapstructure:"max_open_conn"`
	IdleConnections   int32  `mapstructure:"max_idle_conn"`
	MaxConnectionLife int32  `mapstructure:"max_conn_life"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config") // Name of file (without extension)
	viper.SetConfigType("yaml")   // File type
	viper.AddConfigPath("./internal/config/")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err.Error())
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error parsing config file")
	}

	return &config
}
