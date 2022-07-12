package repo

import (
	"fmt"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"Email"`
}

// UserRepo - In a real-world situation, this struct would have types like "db" for SQL operations
type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) AddUser(u User) string {
	return fmt.Sprintf("The user %s with email %s was successfully created", u.Name, u.Email)
}
