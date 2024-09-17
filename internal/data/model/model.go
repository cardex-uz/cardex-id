package model

import (
	"gorm.io/gorm"
	"time"
)

type AuthTrack struct {
	gorm.Model
	phone      string
	verifyCode string
	verifyTime time.Time
	isVerified bool
}

type Profile struct {
	gorm.Model
	username  string
	phone     string
	firstName string
	lastName  string
	avatar    string
}

type Session struct {
	gorm.Model
	userID string
	token  string
}
