package main

import (
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
)

func main() {
	config := config.NewConfig("../.env")

	r := infrastructure.NewRouting(config)

	r.Run(config.ContainerServerPort)
}
