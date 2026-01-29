package domain

import "time"

type Track struct {
	Id int `json:"id"`
	LibraryId int `json:"library_id"`
	Description string `json:"description"`
	UploadedAt time.Time `json:"uploaded_at"`
}