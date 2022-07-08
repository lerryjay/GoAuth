package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"primaryKey;"`
	Email      string    `json:"email" gorm:"uniqueIndex:idx_email"`
	Username   string    `json:"username" gorm:"uniqueIndex:idx_username"`
	Password   string    ``
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Telephone  string    `json:"telephone" gorm:"uniqueIndex:idx_username"`
	Token      int       ``
	Role       int       `json:"role"`
	CreatedAt  time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
	ModifiedAt time.Time `json:"ModifiedAt"  gorm:"autoUpdateTime" `
}

type UserPermissions struct {
	UserID     uint `json:"userid"`
	User       User
	Permission string    `json:"permission" `
	CreatedAt  time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
	ModifiedAt time.Time `json:"ModifiedAt"  gorm:"autoUpdateTime" `
}
