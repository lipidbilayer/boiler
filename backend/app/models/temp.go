package models

type Item struct {
	BaseModel   `bun:"table:orders,alias:or"`
	ID          int64   `bun:",pk,autoincrement" json:"id"`
	Name        string  `bun:",pk"`
	Description string  `json:"description"`
	Orders      []Order `bun:"m2m:order_to_items,join:Item=Order"`
}

type Order struct {
	BaseModel `bun:"table:items,alias:it"`
	ID        int64  `bun:",pk,autoincrement" json:"id"`
	Name      string `json:"name"`
	Menu      string `json:"menu"`
	Type      string `json:"type"`
	Items     []Item `bun:"m2m:order_to_items,join:Order=Item"`
}

type OrderToItem struct {
	OrderID int64  `bun:",pk"`
	Order   *Order `bun:"rel:belongs-to,join:order_id=id"`
	ItemID  int64  `bun:",pk"`
	Item    *Item  `bun:"rel:belongs-to,join:item_id=id"`
}
