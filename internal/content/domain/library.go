package domain

type Library struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	CreatedAt string `json:"created_at"`
	TotalTracks int `json:"total_tracks"`
}