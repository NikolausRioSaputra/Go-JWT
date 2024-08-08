package routes

import (
	"jwtGolang/internal/handler"
	"jwtGolang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, userHandler handler.UserHandlerInterface) {
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	protected := router.Group("/", middleware.AuthMiddleware("secret"))
	protected.GET("/welcome", userHandler.Welcome)
}
