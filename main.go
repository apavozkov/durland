package main

import (
	"durland/config"
	"durland/models"
	"log"
)

func main() {
	if err := config.LoadConfig("config.json"); err != nil {
		log.Fatal(err)
	}

	cfg := config.GetConfig()

	durlian := models.NewDurlian(cfg.Races, cfg.Locations)
	log.Printf("Создан чел: %+v", durlian.KnownInfo)
}
