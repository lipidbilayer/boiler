package dbgobun

import (
	"context"

	"github.com/lipidbilayer/boiler/app/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *DatabaseBun) GetUserWithPassword(ctx context.Context, user *models.User) error {
	err := d.DB.NewSelect().Model(user).ExcludeColumn("password").Where("username = ?username").Where("password = crypt(?, password)", user.Password).Scan(ctx)
	return d.errorDatabase(err, "User")
}

func (d *DatabaseBun) IndexUser(ctx context.Context) ([]*models.User, error) {
	var items []*models.User
	err := d.DB.NewSelect().Model(&items).Scan(ctx)
	for index := range items {
		items[index].Password = nil
	}
	return items, err
}

func (d *DatabaseBun) ShowUser(ctx context.Context, item *models.User) error {
	err := d.DB.NewSelect().Model(item).WherePK().Scan(ctx)
	if item != nil {
		item.Password = nil
	}
	return d.errorDatabase(err, "User")
}

func (d *DatabaseBun) CreateUser(ctx context.Context, item *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*item.Password), bcrypt.DefaultCost)
	if err != nil {
		return d.errorDatabase(err, "User")
	}
	stringHashedPassword := string(hashedPassword)
	item.Password = &stringHashedPassword
	_, err = d.DB.NewInsert().Model(item).Exec(ctx)
	item.Password = nil
	return d.errorDatabase(err, "User")
}

func (d *DatabaseBun) UpdateUser(ctx context.Context, item *models.User) error {
	if item.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*item.Password), bcrypt.DefaultCost)
		if err != nil {
			return d.errorDatabase(err, "User")
		}
		stringHashedPassword := string(hashedPassword)
		item.Password = &stringHashedPassword
	}

	_, err := d.DB.NewUpdate().Model(item).WherePK().OmitZero().Exec(ctx)
	if err != nil {
		return d.errorDatabase(err, "User")
	}

	err = d.ShowUser(ctx, item)
	if err != nil {
		return d.errorDatabase(err, "User")
	}

	return d.errorDatabase(err, "User")
}

func (d *DatabaseBun) DeleteUser(ctx context.Context, item *models.User) error {
	_, err := d.DB.NewDelete().Model(item).WherePK().Exec(ctx)
	if item != nil {
		item.Password = nil
	}
	return d.errorDatabase(err, "User")
}
