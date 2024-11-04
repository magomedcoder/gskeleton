package domain

import "strings"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) IsValidEmail() bool {
	// Простой пример валидации email
	return strings.Contains(u.Email, "@")
}
