package dbgobun

import (
	"context"

	"github.com/lipidbilayer/boiler/app/models"
)

func (d *DatabaseBun) IndexOrder(ctx context.Context) ([]*models.Order, error) {
	var items []*models.Order
	err := d.DB.NewSelect().Model(&items).Relation("Items").Scan(ctx)
	return items, err
}

func (d *DatabaseBun) IndexRole(ctx context.Context) ([]*models.Role, error) {
	var items []*models.Role
	err := d.DB.NewSelect().Model(&items).Relation("Accesses").Scan(ctx)
	return items, err
}
