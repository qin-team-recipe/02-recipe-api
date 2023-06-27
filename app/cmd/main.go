package main

import (
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
)

func main() {
	config := config.NewConfig(".")

	db := infrastructure.NewDB(config)
	google := infrastructure.NewGoogle(config)

	r := infrastructure.NewRouting(config, db, google)

	r.Run(config.ContainerServerPort)
}
