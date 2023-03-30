package main

import (
	"log"

	"github.com/cyberstz/notatka/lib/config"
)

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatalf("Could not load config. %v", err)
	}

	log.Fatalf("Config. %v", config)
}
