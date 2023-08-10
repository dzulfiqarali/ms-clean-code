package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database struct {
		Driver   string `mapstructure:"DRIVER"`
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
		Database string `mapstructure:"DATABASE"`
	} `mapstructure:"DATABASE"`

	Elastic struct {
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
		Index    string `mapstructure:"INDEX"`
	} `mapstructure:"ELASTIC"`

	FakeApi struct {
		Host string `mapstructure:"HOST"`
	} `mapstructure:"FAKEAPI"`

	Service struct {
		Host string `mapstructure:"HOST"`
		Port string `mapstructure:"PORT"`
	} `mapstructure:"SERVICE"`
}

func LoadConfig() (AppConfig *Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Failed to unmarshal configuration: %v", err)
	}

	return
}
