package main

import (
	"context"
	"log"
	"metaverse/internal/infrastructure/database"
	"metaverse/settings"
)

func main() {
	myContext := context.Background()
	myConfig, err := settings.New()
	if err != nil {
		log.Panicf("failed to load settings %s", err)
	}

	_, err = database.New(myContext, myConfig)
	if err != nil {
		log.Panicf("failed to start database %s", err)
	}
}
