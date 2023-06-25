package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/console"
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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func (r *Routing) setRouting() {

	chefsController := product.NewChefsController(r.DB)
	chefFollowsController := product.NewChefFollowsController(r.DB)
	chefRecipesController := product.NewChefRecipesController(r.DB)
	recipeFavoritesController := product.NewRecipeFavoritesController(r.DB)
	recipeIngredientsController := product.NewRecipeIngretientsController(r.DB)
	recipeLinksController := product.NewRecipeLinksController(r.DB)
	recipeStepsController := product.NewRecipeStepsController(r.DB)
	shoppingItemsController := product.NewShoppingItemsController(r.DB)
	userController := product.NewUsersController()
	userRecipesController := product.NewUserRecipesController(r.DB)

	// REST API用
	v1 := r.Gin.Group("/api/v1")
	// swagger用
	docs.SwaggerInfo.BasePath = "/api/v1"
	{
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
		 * chef recipes
		 *
		 */
		v1.GET("/chefRecipes", func(ctx *gin.Context) {
			chefRecipesController.GetList(ctx)
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

		v1.GET("/users", func(ctx *gin.Context) {
			userController.Get(ctx)
		})

		/*
		 * recipes ingredients
		 *
		 */
		v1.POST("/recipeIngredients", func(ctx *gin.Context) {
			recipeIngredientsController.Post(ctx)
		})

		/*
		 * recipes links
		 *
		 */
		v1.POST("/recipeLinks", func(ctx *gin.Context) {
			recipeLinksController.Post(ctx)
		})

		/*
		 * recipes steps
		 *
		 */
		v1.POST("/recipeSteps", func(ctx *gin.Context) {
			recipeStepsController.Post(ctx)
		})

		/*
		 * shopping Items
		 *
		 */
		v1.GET("/shoppingItems", func(ctx *gin.Context) {
			shoppingItemsController.GetList(ctx)
		})
		v1.POST("/shoppingItems", func(ctx *gin.Context) {
			shoppingItemsController.Post(ctx)
		})

		v1.PATCH("/shoppingItems/:id", func(ctx *gin.Context) {
			shoppingItemsController.Patch(ctx)
		})
		v1.DELETE("/shoppingItems/:id", func(ctx *gin.Context) {
			shoppingItemsController.Delete(ctx)
		})

		/*
		 * user recipes
		 *
		 */
		v1.GET("/userRecipes", func(ctx *gin.Context) {
			userRecipesController.GetList(ctx)
		})
		v1.POST("/userRecipes", func(ctx *gin.Context) {
			userRecipesController.Post(ctx)
		})

		/*
		 * swagger
		 *
		 */
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	consoleChefsController := console.NewChefsController(r.DB)
	consoleRecipesController := console.NewRecipesController(r.DB)

	v1Console := r.Gin.Group("/api/v1/console")
	{

		/*
		 * console chefs
		 *
		 */
		v1Console.POST("/chefs", func(ctx *gin.Context) {
			consoleChefsController.Post(ctx)
		})

		/*
		 * console chef recipes
		 *
		 */
		v1Console.POST("/chefRecipes", func(ctx *gin.Context) {
			consoleRecipesController.Post(ctx)
		})
	}
}

func (r *Routing) Run(port string) {
	r.Gin.Run(fmt.Sprintf(":%s", port))
}
