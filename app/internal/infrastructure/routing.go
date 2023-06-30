package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/console"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"

	docs "github.com/qin-team-recipe/02-recipe-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Routing struct {
	Gin    *gin.Engine
	DB     *DB
	Google *Google
	Jwt    token.Maker
}

func NewRouting(c *config.Config, db *DB, google *Google, jwt token.Maker) *Routing {
	r := &Routing{
		DB:     db,
		Gin:    gin.Default(),
		Google: google,
		Jwt:    jwt,
	}

	// r.setCors()

	r.setRouting()

	return r
}

//	@title			Team02's API
//	@version		1.0
//	@description	This is a Team02's API Docs at Qin.
//	@termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/api/v1
func (r *Routing) setRouting() {

	authenticatesController := product.NewAuthenticatesController(r.Google)
	chefsController := product.NewChefsController(r.DB)
	chefFollowsController := product.NewChefFollowsController(r.DB)
	chefRecipesController := product.NewChefRecipesController(r.DB)
	meController := product.NewMeController(product.MeControllerProvider{DB: r.DB, Google: r.Google, Jwt: r.Jwt})
	recipeFavoritesController := product.NewRecipeFavoritesController(r.DB)
	recipeIngredientsController := product.NewRecipeIngretientsController(r.DB)
	recipeLinksController := product.NewRecipeLinksController(r.DB)
	recipeStepsController := product.NewRecipeStepsController(r.DB)
	shoppingItemsController := product.NewShoppingItemsController(r.DB)
	userController := product.NewUsersController(&product.UsersControllerProvider{DB: r.DB, Google: r.Google})
	userRecipesController := product.NewUserRecipesController(r.DB)
	userShoppingItemsController := product.NewUserShoppingItemsController(r.DB)

	// REST API用
	v1 := r.Gin.Group("/api/v1")
	// swagger用
	docs.SwaggerInfo.BasePath = "/api/v1"
	{
		//	@summary		Test API.
		//	@description	This API return 'Hello World!!'.
		//	@tags			mock
		//	@accept			application/x-json-stream
		//	@Success		200	{object}	json
		//	@router			/ [get]
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
		})

		/*
		 * authenticates
		 *
		 */
		v1.GET("/authenticates/google", func(ctx *gin.Context) {
			authenticatesController.GetGoogle(ctx)
		})
		v1.GET("/authenticates/google/userinfo", func(ctx *gin.Context) {
			authenticatesController.GetGoogleUserInfo(ctx)
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
		 * me
		 *
		 */
		v1.GET("/me", func(ctx *gin.Context) {
			meController.Get(ctx)
		})
		v1.GET("/me/login", func(ctx *gin.Context) {
			meController.LoginUser(ctx)
		})

		v1.POST("/me", func(ctx *gin.Context) {
			meController.Post(ctx)
		})

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

		v1.POST("/users", func(ctx *gin.Context) {
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
		 * shopping items
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
		 * user shopping Items
		 *
		 */
		v1.GET("/userShoppingItems", func(ctx *gin.Context) {
			userShoppingItemsController.GetList(ctx)
		})
		v1.POST("/userShoppingItems", func(ctx *gin.Context) {
			userShoppingItemsController.Post(ctx)
		})

		v1.PATCH("/userShoppingItems/:id", func(ctx *gin.Context) {
			userShoppingItemsController.Patch(ctx)
		})
		v1.DELETE("/userShoppingItems/:id", func(ctx *gin.Context) {
			userShoppingItemsController.Delete(ctx)
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
