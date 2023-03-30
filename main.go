package main

import (
	"log"

	"github.com/cyberstz/notatka/lib/config"
	logger "github.com/cyberstz/notatka/lib/logger"
)

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatalf("Could not load config. %v", err)
	}

	l := logger.New(config.LogLevel)

	l.Info().Msg("Test zerologger info message.")
}
