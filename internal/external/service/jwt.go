package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jhamiltonjunior/erp-golang/internal/domain/entities"
	"os"
	"time"
)

type JWT struct{}

func (t *JWT) GenerateToken(id entities.UserID, roles []string, permissions map[string]string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     id,
		"exp":         time.Now().Add(30 * 24 * time.Hour).Unix(),
		"authorized":  true,
		"role":        roles,
		"permissions": permissions,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *JWT) ValidateToken() {

}
