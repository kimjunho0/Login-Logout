package models

import (
	"time"
)

type UserInformation struct {
	ID        string `json:"id"gorm:"primaryKey;autoIncrement;"`
	UserID    string `json:"userID"`
	Password  string `json:"password"`
	CreatedAt time.Time
	DeletedAt *time.Time
	LoginAt   time.Time
	LogoutAt  time.Time
}
