package main

import (
	"go-echo-clean-architecuture/controller"
	"go-echo-clean-architecuture/db"
	"go-echo-clean-architecuture/repository"
	"go-echo-clean-architecuture/router"
	"go-echo-clean-architecuture/usecase"
)

func main() {
	db := db.NewDB()
	// userValidator := validator.NewUserValidator()
	// taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	// taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	// taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
