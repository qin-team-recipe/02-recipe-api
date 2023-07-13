package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/qin-team-recipe/02-recipe-api/example/tasks/utilities"
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

func CreateRecipeLinks(db *infrastructure.DB, benchmark *utilities.Benchmark) (error) {

	interactor := product.RecipeLinkInteractor{
		DB:   &gateways.DBRepository{DB: db},
		Recipe: &repository.RecipeRepository{},
		RecipeLink: &repository.RecipeLinkRepository{},
	}
	
	j, err := ioutil.ReadFile("/app/example/data/example_recipe_links_data.json")
	if err != nil {
		log.Println(err)
		benchmark.Finish()
		return err
	}

	recipeLinks := []domain.RecipeLinks{}

	json.Unmarshal(j, &recipeLinks)
	fmt.Printf("%+v\n", recipeLinks)

	for _, ri := range recipeLinks {
		fmt.Printf("%+v\n", ri)
		_, res := interactor.Create(&ri)
		if res.Error != nil {
			log.Println(res.Error)
			benchmark.Finish()
			return res.Error
		}
	}

	return nil
}