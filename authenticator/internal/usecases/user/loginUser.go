package usecase

import (
	"authenticator/internal/core"
	"authenticator/internal/entity"
	"authenticator/internal/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(request *core.LoginRequest, userRepository entity.IUserRepository) (string, error) {
	user, err := userRepository.FindByPhone(request.Phone)
	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", errors.New("Invalid Credentials")
	}

	token, err := utils.GenerateJWTFromUser(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
