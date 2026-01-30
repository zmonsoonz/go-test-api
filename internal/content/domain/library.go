package domain

import "time"

type Library struct {
	Id int `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}