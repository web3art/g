package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openSqlite(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
