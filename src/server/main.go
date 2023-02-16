package main

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"picket-main-service/src/cmd"
	"picket-main-service/src/config"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Logger = log.With().Caller().Logger()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	err = cmd.GetRoot(cfg).Execute()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
