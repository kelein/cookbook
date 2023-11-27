package jwtool

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"cookbook/devto-grpc/model"
)

// Manager for JWT authentication
type Manager struct {
	secretKey string
	expired   time.Duration
}

// NewManager creates a new JWT Manager
func NewManager(secretKey string, expired time.Duration) *Manager {
	return &Manager{
		secretKey: secretKey,
		expired:   expired,
	}
}

// Generate create a user token
func (m *Manager) Generate(user *model.User) (string, error) {
	claims := UserClaims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.expired).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

// UserClaims stands for user claim
type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
