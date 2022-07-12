package main

import (
	"fmt"

	"github.com/tiagocesar/awesomeProject/internal/repo"
	"github.com/tiagocesar/awesomeProject/internal/user"
)

func main() {
	// Building the dependency graph
	userRepo := repo.NewUserRepo()
	userService := user.NewService(userRepo)

	fmt.Println(userService.AddUser("Me", "me@someone.com"))
}
