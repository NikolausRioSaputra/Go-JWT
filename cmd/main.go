package main

import (
	"jwtGolang/internal/handler"
	"jwtGolang/internal/provider/db"
	"jwtGolang/internal/repository"
	"jwtGolang/internal/routes"
	"jwtGolang/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserusecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Initialize the router
	route := gin.New()
	routes.InitializeRoutes(route, userHandler)

	// Start the server
	if err := route.Run(":8080"); err != nil {
		panic(err)
	}

}
