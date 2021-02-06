package infrastructure

import (
	"authenticator/internal/core"
	"authenticator/internal/modules/user"
	"authenticator/internal/utils"
	"net/http"
	"strings"

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

	v1.GET("/jwt/verify", router.jwtHandler)
	user := v1.Group("/user")
	router.userRoute(user)
}

func (router *Router) userRoute(r *gin.RouterGroup) {
	userController := user.NewUserController(router.repositories)

	r.POST("/register", userController.CreateUser)
	r.POST("/login", userController.LoginUser)
}

func (router *Router) jwtHandler(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   "Authorization header is empty",
		})
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   "Invalid authorization token format",
		})
		return
	}

	token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   claims,
	})
}
