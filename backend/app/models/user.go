package models

type User struct {
	BaseModel `bun:"table:users,alias:u"`
	ID        int64   `bun:",pk,autoincrement" json:"id"`
	Username  string  `json:"username"`
	Password  *string `json:"password,omitempty"`
	Name      string  `json:"name"`
	Role      string  `json:"role"`
}
