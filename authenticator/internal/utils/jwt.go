package utils

import (
	"authenticator/internal/entity"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaim struct {
	jwt.StandardClaims
	Name      string
	Phone     string
	Role      string
	CreatedAt time.Time
}

func GenerateJWTFromUser(user *entity.User) (string, error) {
	claims := CustomClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APP_NAME"),
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Name:      user.Name,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: time.Unix(user.CreatedAt, 0),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("Signing method invalid")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, err
	}

	return claims, nil
}
