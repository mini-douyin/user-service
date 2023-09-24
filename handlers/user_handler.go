package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

type UserWithProfileResponse struct {
	ID             uint   `json:"id"`
	Email          string `json:"email"`
	Avatar         string `json:"avatar"`
	Background     string `json:"background"`
	Signature      string `json:"signature"`
	FollowingCount int    `json:"following_count"`
	FollowerCount  int    `json:"follower_count"`
	LikesGiven     int    `json:"likes_given"`
	LikesReceived  int    `json:"likes_received"`
	VideoCount     int    `json:"video_count"`
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

func (h *UserHandler) GetUserById(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid user ID format. %v", err.Error())})
		return
	}

	user, err := h.service.GetUserWithProfileById(uint(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("User not found. %v", err.Error())})
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Error fetching user. %v", err.Error())})
			return
		}
	}

	response := UserWithProfileResponse{
		ID:             user.ID,
		Email:          user.Email,
		Avatar:         user.Profile.Avatar,
		Background:     user.Profile.Background,
		Signature:      user.Profile.Signature,
		FollowingCount: user.Profile.FollowingCount,
		FollowerCount:  user.Profile.FollowerCount,
		LikesGiven:     user.Profile.LikesGiven,
		LikesReceived:  user.Profile.LikesReceived,
		VideoCount:     user.Profile.VideoCount,
	}

	c.JSON(http.StatusOK, response)
}
