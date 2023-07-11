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
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/console"
)

func CreateRecipes(db *infrastructure.DB, benchmark *utilities.Benchmark) (error) {

	ri := console.RecipeInteractor{
		DB:   &gateways.DBRepository{DB: db},
		ChefRecipe: &repository.ChefRecipeRepository{},
		Recipe: &repository.RecipeRepository{},
	}
	
	rj, err := ioutil.ReadFile("/app/example/data/example_recipes_data.json")
	if err != nil {
		log.Println(err)
		benchmark.Finish()
		return err
	}
	crj, err := ioutil.ReadFile("/app/example/data/example_chef_recipes_data.json")
	if err != nil {
		log.Println(err)
		benchmark.Finish()
		return err
	}

	recipes := []domain.Recipes{}
	chefRecipes := []domain.ChefRecipes{}

	json.Unmarshal(rj, &recipes)
	json.Unmarshal(crj, &chefRecipes)

	for i, r := range recipes {
		fmt.Printf("%+v\n", r)
		_, res := ri.Create(chefRecipes[i].ChefID, &r)
		if res.Error != nil {
			log.Println(res.Error)
			benchmark.Finish()
			return res.Error
		}
	}

	return nil
}