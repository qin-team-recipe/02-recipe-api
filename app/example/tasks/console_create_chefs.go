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

func CreateChefs(db *infrastructure.DB, benchmark *utilities.Benchmark) (error) {

	interactor := console.ChefInteractor{
		DB:   &gateways.DBRepository{DB: db},
		Chef: &repository.ChefRepository{},
	}
	
	j, err := ioutil.ReadFile("/app/example/data/example_chefs_data.json")
	if err != nil {
		log.Println(err)
		benchmark.Finish()
		return err
	}

	chefs := []domain.Chefs{}

	json.Unmarshal(j, &chefs)
	fmt.Printf("%+v\n", chefs)

	for _, c := range chefs {
		fmt.Printf("%+v\n", c)
		_, res := interactor.Create(&c)
		if res.Error != nil {
			log.Println(res.Error)
			benchmark.Finish()
			return res.Error
		}

	}

	return nil
}