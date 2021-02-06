package core

import (
	"authenticator/internal/entity"
	userRepository "authenticator/internal/repositories/user"
	"database/sql"
)

type Repository struct {
	entity.IUserRepository
}

func InitRepository(db *sql.DB) *Repository {
	return &Repository{
		IUserRepository: userRepository.NewUserRepository(db),
	}
}
