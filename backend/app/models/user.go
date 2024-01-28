package models

import "time"

type User struct {
	BaseModel   `bun:"table:users,alias:us"`
	ID          int64      `bun:",pk,autoincrement" json:"id"`
	Username    string     `json:"username"`
	Password    *string    `json:"password,omitempty"`
	Name        string     `json:"name"`
	Role        string     `json:"role"`
	RoleDetail  *Role      `bun:"rel:belongs-to,join:role=name"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}
