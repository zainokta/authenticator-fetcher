package infrastructure

import (
	"authenticator/internal/core"
	"authenticator/internal/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	repositories *core.Repository
}

func NewRouter(repositories *core.Repository) *Router {
	return &Router{repositories: repositories}
}

func (router *Router) SetRoutes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome!",
		})
	})

	api := route.Group("/api")

	v1 := api.Group("/v1")

	user := v1.Group("/user")
	router.userRoute(user)
}

func (router *Router) userRoute(r *gin.RouterGroup) {
	userController := user.NewUserController(router.repositories)

	r.POST("/register", userController.CreateUser)
}
