package domain

import "time"

type User struct {
	Id int `json:"id" db:"id"`
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	Password string `json:"password" binding:"required"  db:"password_hash"`
}