package main

import (
	"example_clean/controller"
	// "example_clean/db"
	"example_clean/interactor"
	"example_clean/presenter"
	"example_clean/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// dbConn, err := db.NewConnection()
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	userRepository := repository.NewUserRepository()
	userInteractor := interactor.NewUserInteractor(userRepository)
	userPresenter := presenter.NewUserPresenter()
	userController := controller.NewUserController(userInteractor, userPresenter)

	r.GET("/users/:id", userController.GetUser)

	r.Run()
}
