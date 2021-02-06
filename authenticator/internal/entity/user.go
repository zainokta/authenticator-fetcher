package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//User is user entity
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at"`
}

type IUserRepository interface {
	Create(name, phone, role string) (string, error)
}

func ValidateUserName(name string) error {
	if name == "" {
		return errors.New("Name cannot be empty")
	}

	if len(name) < 4 {
		return errors.New("Name must be more than 4 characters")
	}

	return nil
}

func ValidateUserRole(role string) error {
	if role == "" {
		return errors.New("Role cannot be empty, role must be either admin or user")
	}

	if role != "admin" && role != "user" {
		return errors.New("Role not found, role must be either admin or user")
	}

	return nil
}

func ValidateUserPhone(phone string) error {
	if phone == "" {
		return errors.New("Phone number cannot be empty")
	}

	re := regexp.MustCompile("(0|\\+62|062|62)[0-9]+$")

	if !re.MatchString(phone) || len(phone) < 10 || len(phone) > 13 {
		return errors.New("Phone number is not valid")
	}

	return nil
}

func ValidateUser(name, phone, role string) error {
	err := ValidateUserName(name)
	if err != nil {
		return err
	}

	err = ValidateUserPhone(phone)
	if err != nil {
		return err
	}

	err = ValidateUserRole(role)
	if err != nil {
		return err
	}

	return nil
}

func HashUserPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 5)
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}
