package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	UserQSvcUrl string `mapstructure:"USER_QUERY_SVC_URL"`
	UserCSvcUrl string `mapstructure:"USER_COMMAND_SVC_URL"`
	IsDOcker    string `mapstructure:"IS_DOCKER"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")

	path := ".env"

	if os.Getenv("IS_DOCKER") == "true" || os.Getenv("IS_DOCKER") == "1" {
		path = ".docker.env"
	}

	viper.SetConfigName(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
