package gorad

import (
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
)

func OpenSQLX(driver string, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.MapperFunc(strcase.ToSnake)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
