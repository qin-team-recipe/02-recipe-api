package main

import (
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/example/tasks"
	"github.com/qin-team-recipe/02-recipe-api/example/tasks/utilities"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
)

func main() {
	benchmark := new(utilities.Benchmark)
	benchmark.Start()

	config := config.NewConfig(".")
	db := infrastructure.NewDB(config)

	// chefs id:   [10001, 10002]
	err := tasks.CreateChefs(db, benchmark)
	if err != nil {
		return
	}

	// recipes id: [20001, 20002, 20003]
	err = tasks.CreateRecipes(db, benchmark)
	if err != nil {
		return
	}

	err = tasks.CreateRecipeIngredients(db, benchmark)
	if err != nil {
		return
	}

	err = tasks.CreateRecipeSteps(db, benchmark)
	if err != nil {
		return
	}

	err = tasks.CreateRecipeLinks(db, benchmark)
	if err != nil {
		return
	}

	benchmark.Finish()
}
