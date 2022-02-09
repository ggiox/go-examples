package main

import (
	"fmt"
	"log"

	"ApiGORM/application/repositories"
	"ApiGORM/domain"
	"ApiGORM/framework/utils"
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
