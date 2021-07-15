package config

import (
	"database/sql"
)

type Config struct {
	setting Setting
	db      *sql.DB
}

func NewConfig(
	setting Setting,
	db *sql.DB,
) Config {
	return Config{
		setting: setting,
		db:      db,
	}
}
