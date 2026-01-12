package domain

type Track struct {
	Id int `json:"id"`
	LibraryId int `json:"library_id"`
	Duration int `json:"duration"`
	FilePath string `json:"file_path"`
	UploadedAt string `json:"uploaded_at"`
}