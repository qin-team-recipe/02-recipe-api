package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/product"

	docs "github.com/qin-team-recipe/02-recipe-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	chefFollowsController := product.NewChefFollowsController(r.DB)
	recipesController := product.NewRecipesController(r.DB)
	recipeFavoritesController := product.NewRecipeFavoritesController(r.DB)
	userController := product.NewUsersController()

	// REST API用
	v1 := r.Gin.Group("/v1")
	// swagger用
	docs.SwaggerInfo.BasePath = "/v1"

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

	v1.GET("/chefs/:screenName", func(ctx *gin.Context) {
		chefsController.Get(ctx)
	})

	/*
	 * chefs　follows
	 *
	 */
	v1.GET("/chefFollows", func(ctx *gin.Context) {
		chefFollowsController.GetList(ctx)
	})

	/*
	 * recipes
	 *
	 */
	v1.GET("/recipes", func(ctx *gin.Context) {
		recipesController.GetList(ctx)
	})

	// v1.GET("/recipes/:id", func(ctx *gin.Context) {
	// 	recipesController.Get(ctx)
	// })

	/*
	 * recipes favorites
	 *
	 */
	v1.GET("/recipeFavorites", func(ctx *gin.Context) {
		recipeFavoritesController.GetList(ctx)
	})

	/*
	 * users
	 *
	 */
	v1.GET("/users", func(ctx *gin.Context) {
		userController.Get(ctx)
	})

	/*
	 * swagger
	 *
	 */
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (r *Routing) Run(port string) {
	r.Gin.Run(fmt.Sprintf(":%s", port))
}
