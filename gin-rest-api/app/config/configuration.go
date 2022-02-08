package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

func SetupConfig(configPath string) Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)

	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("server for this application is %s", configuration.Server.Port)

	return configuration
}
