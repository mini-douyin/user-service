package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/models"
	"user-service/services"
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

	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered!",
	})
}
