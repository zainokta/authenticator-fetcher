package usecase

import (
	"authenticator/internal/core"
	"authenticator/internal/entity"
)

func CreateUser(request *core.UserRequest, userRepository entity.IUserRepository) (string, error) {
	err := entity.ValidateUser(request.Name, request.Phone, request.Role)
	if err != nil {
		return "", err
	}

	password, err := userRepository.Create(request.Name, request.Phone, request.Role)
	if err != nil {
		return "", err
	}

	return password, nil
}
