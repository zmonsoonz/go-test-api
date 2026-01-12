package domain

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	CreatedAt  string `json:"created_at"`
	Password string `json:"password"`
}