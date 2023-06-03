package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qin-team-recipe/02-recipe-api/config"
)

type Routing struct {
	Gin *gin.Engine
}

func NewRouting(c *config.Config) *Routing {
	r := &Routing{
		Gin: gin.Default(),
	}

	// r.setCors()

	r.setRouting()

	return r
}

func (r *Routing) setRouting() {
	// REST API用
	v1 := r.Gin.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
	})
}

func (r *Routing) Run(port string) {
	r.Gin.Run(fmt.Sprintf(":%s", port))
}
