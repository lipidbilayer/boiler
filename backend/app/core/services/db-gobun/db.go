package dbgobun

import (
	"database/sql"
	"log"

	service "github.com/lipidbilayer/boiler/app/core/services"
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
	return &dbService
}

func (d *DatabaseBun) Stop() {
	if err := db.Close(); err != nil {
		log.Panic("Failed to close the database", "error", err)
	}
}
