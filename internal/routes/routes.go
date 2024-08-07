package routes

import (
	"jwtGolang/internal/handler"
	"jwtGolang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	protected := r.Group("/", middleware.AuthMiddleware("secret"))
	protected.GET("/welcome", userHandler.Welcome)

	return r
}
