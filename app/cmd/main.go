package main

import (
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

func main() {
	config := config.NewConfig(".")
	jwt := token.NewJwtMaker(config)

	db := infrastructure.NewDB(config)
	google := infrastructure.NewGoogle(config)

	r := infrastructure.NewRouting(config, db, google, jwt)

	r.Run(config.ContainerServerPort)
}
