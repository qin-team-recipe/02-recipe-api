package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/example/tasks/utilities"
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/console"
)

func main() {
	benchmark := new(utilities.Benchmark)
	benchmark.Start()

	config := config.NewConfig(".")
	db := infrastructure.NewDB(config)

	interactor := getIntaractor(db)

	/*
		ここに模擬データを読み込ませる処理を書く
	*/
	j, err := ioutil.ReadFile("./example/tasks/example_chef_data.json")
	if err != nil {
		log.Println(err)
		benchmark.Finish()
		return
	}

	c := &domain.Chefs{}

	json.Unmarshal(j, c)
	fmt.Printf("%+v", c)

	_, res := interactor.Create(c)
	if res.Error != nil {
		log.Println(res.Error)
		benchmark.Finish()
		return
	}
	/*
		ここまでが一例
		配列の場合はfor文などで一つずつ処理をする
	*/

	benchmark.Finish()
}

// interactor のインスタンス化
func getIntaractor(db gateways.DB) console.ChefInteractor {
	return console.ChefInteractor{
		DB:   &gateways.DBRepository{DB: db},
		Chef: &repository.ChefRepository{},
	}
}
