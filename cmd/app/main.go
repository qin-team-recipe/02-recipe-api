package main

import (
	"log"

	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
)

func main() {
	config := config.NewConfig("../.env")
	log.Println(config)

	r := infrastructure.NewRouting(config)

	r.Run(config.ServerPort)
}
