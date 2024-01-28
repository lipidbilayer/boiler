package dbgobun

import (
	"database/sql"
	"fmt"
	"log"

	service "github.com/lipidbilayer/boiler/app/core/services"
	"github.com/lipidbilayer/boiler/app/models"
	"github.com/lipidbilayer/boiler/lib/apperror"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

type DatabaseBun struct {
	DB *bun.DB
	// timeloc *time.Location
}

func New(config service.ConfigService) *DatabaseBun {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.GetDatabaseURL())))

	db = bun.NewDB(sqldb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.WithEnabled(config.GetDebugMode())))

	dbService := DatabaseBun{db}
	dbService.RegisterTable()
	return &dbService
}

func (d *DatabaseBun) RegisterTable() {
	db.RegisterModel((*models.RoleToAccess)(nil))
	db.RegisterModel((*models.OrderToItem)(nil))
}

func (d *DatabaseBun) Stop() {
	if err := db.Close(); err != nil {
		log.Panic("Failed to close the database", "error", err)
	}
}

func (d *DatabaseBun) errorDatabase(err error, model string) error {
	if err == nil {
		return nil
	}
	switch err.Error() {
	case "sql: no rows in result set":
		return apperror.NewError(err, fmt.Sprintf("%s tidak di temukan", model), apperror.NotFoundError)
	}
	// switch err {
	// case pg.ErrNoRows:
	// 	return apperror.NewError(err, fmt.Sprintf("%s tidak di temukan", model), apperror.NotFoundError)
	// }
	return err
}
