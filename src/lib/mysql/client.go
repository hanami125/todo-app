package mysql

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // import mysql
	"github.com/jmoiron/sqlx"
)

// DBConfig ... set db config
type DBConfig struct {
	ConnectionName string
	User           string
	Password       string
	Database       string
}

// NewSQLClient ... get new sql client
func NewSQLClient(ctx context.Context, cfg *DBConfig) (*sqlx.DB, error) {
	ds := fmt.Sprintf("%s:%s@/%s",
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	db, err := sqlx.Open("mysql", ds)
	if err != nil {
		return nil, fmt.Errorf("mysql: fails to open connection: %w", err)
	}

	return db, nil
}
