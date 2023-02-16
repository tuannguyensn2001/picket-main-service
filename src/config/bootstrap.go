package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

type structure struct {
	AppHttpPort              string `mapstructure:"APP_HTTP_PORT"`
	AppGrpcPort              string `mapstructure:"APP_GRPC_PORT"`
	AppHttpAddress           string `mapstructure:"APP_HTTP_ADDRESS"`
	AppEnv                   string `mapstructure:"APP_ENV"`
	AppSecretKey             string `mapstructure:"APP_SECRET_KEY"`
	Oauth2GoogleClientId     string `mapstructure:"OAUTH2_GOOGLE_CLIENT_ID"`
	Oauth2GoogleClientSecret string `mapstructure:"OAUTH2_GOOGLE_CLIENT_SECRET"`
	ClientUrl                string `mapstructure:"CLIENT_URL"`
	DatabaseUrl              string `mapstructure:"DATABASE_URL"`
}

func bootstrap() structure {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Send()
	}
	var result structure
	if err := viper.Unmarshal(&result); err != nil {
		log.Fatal().Err(err).Send()
	}

	return result
}
