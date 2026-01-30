package domain

type User struct {
	Id int `json:"id" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Email string `json:"email" binding:"required" db:"email"`
	CreatedAt  int64 `json:"created_at" db:"created_at"`
	Password string `json:"password" binding:"required"  db:"password_hash"`
}