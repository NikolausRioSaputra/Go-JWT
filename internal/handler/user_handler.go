package handler

import (
	"jwtGolang/internal/domain"
	"jwtGolang/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Welcome(c *gin.Context)
}
type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler (userUsecase usecase.UserUsecase) UserHandlerInterface {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var u domain.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.UserUsecase.Register(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var u domain.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.UserUsecase.Login(c.Request.Context(), u)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) Welcome(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + user.(string)})
}
