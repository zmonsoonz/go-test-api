package domain

type Track struct {
	Id int `json:"id" db:"id"`
	LibraryId int `json:"library_id" db:"library_id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UploadedAt int64 `json:"uploaded_at" db:"uploaded_at"`
}

type UpdateTrackInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
