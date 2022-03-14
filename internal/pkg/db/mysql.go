package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openMysql(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
