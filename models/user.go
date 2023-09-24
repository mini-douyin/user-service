package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint        `gorm:"primaryKey" json:"id"`
	Email    string      `gorm:"uniqueIndex;not null" binding:"required,email,max=100" json:"email"`
	Password string      `gorm:"not null" binding:"required,min=6,max=50" json:"password"`
	Profile  UserProfile `gorm:"foreignKey:UserId" json:"profile"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	userProfile := UserProfile{UserId: user.ID}
	if err = tx.Create(&userProfile).Error; err != nil {
		return err
	}
	return
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

type UserProfile struct {
	gorm.Model
	UserId         uint   `gorm:"unique" json:"user_id"`
	Avatar         string `json:"avatar"`
	Background     string `json:"background"`
	Signature      string `json:"signature"`
	FollowingCount int    `gorm:"default:0" json:"following_count"`
	FollowerCount  int    `gorm:"default:0" json:"follower_count"`
	LikesGiven     int    `gorm:"default:0" json:"likes_given"`
	LikesReceived  int    `gorm:"default:0" json:"likes_received"`
	VideoCount     int    `gorm:"default:0" json:"video_count"`
}
