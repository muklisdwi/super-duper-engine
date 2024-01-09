package config

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// reference : https://pkg.go.dev/github.com/spf13/viper#section-readme
// create structure config
type Config struct {
	Server struct {
		Env      string `mapstructure:"ENV"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
		Port     string `mapstructure:"PORT"`
	} `mapstructure:"SERVER"`
}

const (
	ENV_NAME = ".env"
)

var (
	config Config
	once   sync.Once
)

func InitializeConfiguration() *Config {
	viper.SetConfigFile(ENV_NAME)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Fail reading config file")
	}

	once.Do(func() {
		err = viper.Unmarshal(&config)
		if err != nil {
			log.Fatal().Err(err).Msg("Configuration failed to initialize")
		}
		log.Info().Msg("Configuration initialized")
	})

	return &config
}
