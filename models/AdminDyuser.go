package models

import "golang.org/x/crypto/bcrypt"

type AdminDyUser struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password []byte `json:"-"`
}

// Hash password func
func (user *AdminDyUser) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

// Compare passowrd func
func (user *AdminDyUser) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
