package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	cfg := Config{
		Database: DatabaseConfig{
			Host:     "103.197.188.89",
			Port:     "5432",
			User:     "vansteve123",
			Password: "3225501",
			DBName:   "test",
			SSLMode:  "disable",
			TimeZone: "Asia/Jakarta",
		},
		Server: ServerConfig{
			Port: "8080",
		},
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}
	return &cfg
}
