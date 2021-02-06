package user

import (
	"authenticator/internal/core"
	"authenticator/internal/entity"
	usecase "authenticator/internal/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository entity.IUserRepository
}

func NewUserController(userRepository entity.IUserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}

func (u *UserController) CreateUser(c *gin.Context) {
	request := &core.UserRequest{}
	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	password, err := usecase.CreateUser(request, u.userRepository)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   password,
	})
}

func (u *UserController) LoginUser(c *gin.Context) {
	request := &core.LoginRequest{}
	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	token, err := usecase.LoginUser(request, u.userRepository)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   token,
	})
}
