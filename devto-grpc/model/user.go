package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User represents request auth info
type User struct {
	Username string
	Password string
	Role     string
}

// NewUser creates a new User instance with username and password
func NewUser(username, password, role string) (*User, error) {
	cryptPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to generate crypto password: %v", err)
	}
	return &User{
		Role:     role,
		Username: username,
		Password: string(cryptPass),
	}, nil
}

// Authed check whether the password is correct
func (u *User) Authed(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Clone returns a clone of this user
func (u *User) Clone() *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}
