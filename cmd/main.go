package main

import (
	"jwtGolang/internal/handler"
	"jwtGolang/internal/provider/db"
	"jwtGolang/internal/repository"
	"jwtGolang/internal/routes"
	"jwtGolang/internal/usecase"
)

func main() {
	database := db.InitDB()
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userUsecase := &usecase.UserUsecase{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserUsecase: userUsecase}

	r := routes.SetupRouter(userHandler)
	r.Run(":8080")
}
