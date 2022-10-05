package gorad

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
	gofs "io/fs"
)

type Migration struct {
	logger migrate.Logger
}

type MigrationOption func(*Migration)

func MigrateUp(
	db *sql.DB,
	fs gofs.FS,
	path string,
	opts ...MigrationOption,
) error {
	var m Migration
	for _, opt := range opts {
		opt(&m)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return err
	}

	d, err := iofs.New(fs, path)
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("iofs", d, "sqlite3", driver)
	if err != nil {
		return err
	}
	if m.logger != nil {
		migrator.Log = m.logger
	}

	return migrator.Up()
}

func WithMigrationLogger(logger *zap.Logger) MigrationOption {
	return func(o *Migration) {
		o.logger = newMigrateZapAdapter(logger)
	}
}

// migrateZapAdapter is a wrapper around zap.Logger to satisfy the migrate.Logger interface.
type migrateZapAdapter struct {
	logger *zap.Logger
}

func newMigrateZapAdapter(logger *zap.Logger) *migrateZapAdapter {
	return &migrateZapAdapter{logger: logger}
}

func (ad *migrateZapAdapter) Printf(format string, v ...interface{}) {
	ad.logger.Info(fmt.Sprintf(format, v...))
}

func (*migrateZapAdapter) Verbose() bool {
	return true
}
