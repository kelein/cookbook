package jwtool

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

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

// Verify parse user claims from token
func (m *Manager) Verify(accessToken string) (*UserClaims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected token signed method")
		}
		return []byte(m.secretKey), nil
	}

	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

// UserClaims stands for user claim
type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
