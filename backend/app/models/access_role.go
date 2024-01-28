package models

type Role struct {
	BaseModel   `bun:"table:roles,alias:ro"`
	Name        string   `bun:",pk"`
	Description string   `json:"description"`
	Accesses    []Access `bun:"m2m:role_to_accesses,join:Role=Access"`
}

type Access struct {
	BaseModel `bun:"table:accesses,alias:ac"`
	ID        int64  `bun:",pk,autoincrement" json:"id"`
	Name      string `json:"name"`
	Menu      string `json:"menu"`
	Type      string `json:"type"`
	Roles     []Role `bun:"m2m:role_to_accesses,join:Access=Role"`
}

type RoleToAccess struct {
	RoleName string  `bun:",pk"`
	Role     *Role   `bun:"rel:belongs-to,join:role_name=name"`
	AccessID int64   `bun:",pk"`
	Access   *Access `bun:"rel:belongs-to,join:access_id=id"`
}
