package models

type User struct {
	BaseModel
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Password *string `json:"password,omitempty"`
	Name     string  `json:"name"`
	Role     string  `json:"role"`
}
