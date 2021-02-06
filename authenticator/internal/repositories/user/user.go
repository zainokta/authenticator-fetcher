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

func (u *UserRepository) FindByPhone(phone string) (*entity.User, error) {
	row, err := u.db.Query("SELECT * FROM users WHERE phone = $1", phone)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	user := &entity.User{}
	hasResult := false

	for row.Next() {
		err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Role, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		hasResult = true
	}
	if !hasResult {
		return nil, errors.New("User with phone " + phone + " not found.")
	}

	return user, nil
}
