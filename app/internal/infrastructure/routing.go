package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/infrastructure/middleware"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/console"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"

	docs "github.com/qin-team-recipe/02-recipe-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const basePath = "/api/v1"

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

	r.setCors(middleware.NewCors(c))

	r.setRouting()

	return r
}

// Cors 対応
func (r *Routing) setCors(cors *middleware.Cors) {
	r.Gin.Use(cors.Config)
}

//	@title			Team02's API
//	@version		1.0
//	@description	This is a Team02's API Docs at Qin.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/api/v1
func (r *Routing) setRouting() {

	authenticatesController := product.NewAuthenticatesController(r.Google)
	chefsController := product.NewChefsController(product.ChefsControllerProvider{DB: r.DB, Jwt: r.Jwt})
	chefFollowsController := product.NewChefFollowsController(r.DB)
	chefRecipesController := product.NewChefRecipesController(r.DB)
	meController := product.NewMeController(product.MeControllerProvider{DB: r.DB, Google: r.Google, Jwt: r.Jwt})
	publishStatusesController := product.NewPublishStatusesController(r.DB)
	recommendsController := product.NewRecommendsController(r.DB)
	recipesController := product.NewRecipesController(r.DB)
	recipeFavoritesController := product.NewRecipeFavoritesController(r.DB)
	recipeIngredientsController := product.NewRecipeIngretientsController(r.DB)
	recipeLinksController := product.NewRecipeLinksController(r.DB)
	recipeStepsController := product.NewRecipeStepsController(r.DB)
	shoppingItemsController := product.NewShoppingItemsController(r.DB)
	userController := product.NewUsersController(&product.UsersControllerProvider{DB: r.DB, Google: r.Google})
	userRecipesController := product.NewUserRecipesController(r.DB)
	userShoppingItemsController := product.NewUserShoppingItemsController(r.DB)

	// REST API用
	v1 := r.Gin.Group(basePath)

	// v1Auth := v1.Use(middleware.JwtAuthMiddleware(r.Jwt))
	// swagger用
	docs.SwaggerInfo.BasePath = basePath
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

		// /*
		//  * chefs　follows
		//  *
		//  */
		// v1.GET("/chefFollows", func(ctx *gin.Context) {
		// 	chefFollowsController.GetList(ctx)
		// })

		// v1.POST("/chefFollows", func(ctx *gin.Context) {
		// 	chefFollowsController.Post(ctx)
		// })
		// v1.DELETE("/chefFollows", func(ctx *gin.Context) {
		// 	chefFollowsController.Delete(ctx)
		// })

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

		// /*
		//  * me
		//  *
		//  */
		// v1.GET("/me", func(ctx *gin.Context) {
		// 	meController.Get(ctx)
		// })
		v1.GET("/login", func(ctx *gin.Context) {
			meController.LoginUser(ctx)
		})
		v1.POST("/register", func(ctx *gin.Context) {
			meController.Post(ctx)
		})

		// v1.PATCH("/me", func(ctx *gin.Context) {
		// 	meController.Patch(ctx)
		// })
		// v1.DELETE("/me", func(ctx *gin.Context) {
		// 	meController.Delete(ctx)
		// })

		/*
		 * recommend chefs or recipes
		 *
		 */
		v1.GET("/recommends/chefs", func(ctx *gin.Context) {
			recommendsController.GetRecommendChefList(ctx)
		})
		v1.GET("/recommends/recipes", func(ctx *gin.Context) {
			recommendsController.GetRecommendRecipeList(ctx)
		})

		/*
		 * recipes
		 *
		 */
		v1.GET("/recipes", func(ctx *gin.Context) {
			recipesController.GetList(ctx, r.Jwt)
		})

		v1.GET("/recipes/:watchID", func(ctx *gin.Context) {
			recipesController.Get(ctx)
		})

		// /*
		//  * recipes favorites
		//  *
		//  */
		// v1.GET("/recipeFavorites", func(ctx *gin.Context) {
		// 	recipeFavoritesController.GetList(ctx)
		// })

		// v1.POST("/recipeFavorites", func(ctx *gin.Context) {
		// 	recipeFavoritesController.Post(ctx)
		// })
		// v1.DELETE("/recipeFavorites", func(ctx *gin.Context) {
		// 	recipeFavoritesController.Delete(ctx)
		// })

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
		v1.GET("/recipeIngredients", func(ctx *gin.Context) {
			recipeIngredientsController.GetList(ctx)
		})
		// v1.POST("/recipeIngredients", func(ctx *gin.Context) {
		// 	recipeIngredientsController.Post(ctx)
		// })

		// /*
		//  * recipes links
		//  *
		//  */
		// v1.POST("/recipeLinks", func(ctx *gin.Context) {
		// 	recipeLinksController.Post(ctx)
		// })

		/*
		 * recipes steps
		 *
		 */
		v1.GET("/recipeSteps", func(ctx *gin.Context) {
			recipeStepsController.GetList(ctx)
		})
		// v1.POST("/recipeSteps", func(ctx *gin.Context) {
		// 	recipeStepsController.Post(ctx)
		// })

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

		// /*
		//  * user recipes
		//  *
		//  */
		// v1.GET("/userRecipes", func(ctx *gin.Context) {
		// 	userRecipesController.GetList(ctx)
		// })
		// v1.POST("/userRecipes", func(ctx *gin.Context) {
		// 	userRecipesController.Post(ctx)
		// })

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

	v1Auth := v1.Use(middleware.JwtAuthMiddleware(r.Jwt))
	{
		/*
		 * me
		 *
		 */
		v1Auth.GET("/me", func(ctx *gin.Context) {
			meController.Get(ctx)
		})
		v1Auth.PATCH("/me", func(ctx *gin.Context) {
			meController.Patch(ctx)
		})
		v1Auth.DELETE("/me", func(ctx *gin.Context) {
			meController.Delete(ctx)
		})

		/*
		 * chefs　follows
		 *
		 */
		v1Auth.GET("/chefFollows", func(ctx *gin.Context) {
			chefFollowsController.GetList(ctx)
		})

		v1Auth.POST("/chefFollows", func(ctx *gin.Context) {
			chefFollowsController.Post(ctx)
		})
		v1Auth.DELETE("/chefFollows", func(ctx *gin.Context) {
			chefFollowsController.Delete(ctx)
		})

		/*
		 * publish statuses
		 *
		 */
		v1Auth.PATCH("/publishStatuses", func(ctx *gin.Context) {
			publishStatusesController.Patch(ctx)
		})

		/*
		 * recipes favorites
		 *
		 */
		v1Auth.GET("/recipeFavorites", func(ctx *gin.Context) {
			recipeFavoritesController.GetList(ctx)
		})

		v1Auth.POST("/recipeFavorites", func(ctx *gin.Context) {
			recipeFavoritesController.Post(ctx)
		})
		v1Auth.DELETE("/recipeFavorites", func(ctx *gin.Context) {
			recipeFavoritesController.Delete(ctx)
		})

		/*
		 * recipes ingredients
		 *
		 */
		v1Auth.POST("/recipeIngredients", func(ctx *gin.Context) {
			recipeIngredientsController.Post(ctx)
		})

		/*
		 * recipes links
		 *
		 */
		v1Auth.POST("/recipeLinks", func(ctx *gin.Context) {
			recipeLinksController.Post(ctx)
		})

		/*
		 * recipes steps
		 *
		 */
		v1Auth.POST("/recipeSteps", func(ctx *gin.Context) {
			recipeStepsController.Post(ctx)
		})

		/*
		 * user recipes
		 *
		 */
		v1Auth.GET("/userRecipes", func(ctx *gin.Context) {
			userRecipesController.GetList(ctx)
		})
		v1Auth.POST("/userRecipes", func(ctx *gin.Context) {
			userRecipesController.Post(ctx)
		})

		v1Auth.GET("/userRecipes/:id", func(ctx *gin.Context) {
			userRecipesController.Get(ctx)
		})
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
