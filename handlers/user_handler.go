package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/models"
	"user-service/services"
	"user-service/utils/token"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(svc services.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := token.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to generate token. %v", err.Error())})
		return
	}

	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to register user. %v", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered!",
		"token":   tokenString,
		"user":    user.Email,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var inputData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetUserByEmail(inputData.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password."})
		return
	}

	if !user.CheckPassword(inputData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password."})
		return
	}

	tokenString, err := token.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to generate token. %v", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in!",
		"token":   tokenString,
		"user":    user.Email,
	})
}
