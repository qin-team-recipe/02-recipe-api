package main

import (
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
)

func main() {
	config := config.NewConfig("../.env")

	db := infrastructure.NewDB(config)
	r := infrastructure.NewRouting(config, db)

	r.Run(config.ContainerServerPort)
}
