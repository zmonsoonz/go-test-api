package domain

type Track struct {
	Id int `json:"id"`
	LibraryId int `json:"library_id"`
	Description string `json:"description"`
	UploadedAt int64 `json:"uploaded_at"`
}