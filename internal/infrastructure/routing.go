package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/product"
)

type Routing struct {
	Gin *gin.Engine
	DB  *DB
}

func NewRouting(c *config.Config, db *DB) *Routing {
	r := &Routing{
		DB:  db,
		Gin: gin.Default(),
	}

	// r.setCors()

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {

	chefsController := product.NewChefsController(r.DB)
	userController := product.NewUsersController()
	// REST API用
	v1 := r.Gin.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
	})

	/*
	 * chefs
	 *
	 */
	v1.GET("/chefs", func(ctx *gin.Context) {
		chefsController.GetList(ctx)
	})

	v1.GET("/chefs/:id", func(ctx *gin.Context) {
		chefsController.GetList(ctx)
	})

	/*
	 * users
	 *
	 */
	v1.GET("/users", func(ctx *gin.Context) {
		userController.Get(ctx)
	})
}

func (r *Routing) Run(port string) {
	r.Gin.Run(fmt.Sprintf(":%s", port))
}
