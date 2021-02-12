package main

import (
	"fmt"
	"log"

	"github.com/ggiox/go-examples/ApiGORM/applicatiion/repositories"
	"github.com/ggiox/go-examples/ApiGORM/domain"
	"github.com/ggiox/go-examples/ApiGORM/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Giovanni",
		Email:    "giovanni.aredes@ggiox.com",
		Password: "Start123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)
	if err != nil {
		log.Fatalf("Error during insert the user: %v", err)
	}
	fmt.Println(result)

}
