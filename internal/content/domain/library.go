package domain

import "time"

type Library struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}