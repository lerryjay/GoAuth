package models

import "time"

type User struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Password   string    ``
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Telephone  string    `json:"telephone"`
	Token      int       ``
	Role       int       `json:"role"`
	CreatedAt  time.Time `json:"CreatedAt"`
	ModifiedAt time.Time `json:"ModifiedAt"`
}
