package repositories

import (
	"authenticator/internal/entity"
	"authenticator/internal/utils"
	"database/sql"
	"errors"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(name, phone, role string) (string, error) {
	stmt, err := u.db.Prepare("INSERT INTO users(id, name, phone, role, password, created_at) VALUES($1, $2, $3, $4, $5, $6) ")
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	password := utils.GenerateRandomString(4)
	hashedPassword, err := entity.HashUserPassword(password)
	uuid := entity.GenerateUUID()
	if uuid == "" {
		return "", errors.New("Failed to generate UUID")
	}

	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(uuid, name, phone, role, string(hashedPassword), time.Now().Unix())
	if err != nil {
		return "", err
	}

	return password, nil
}
