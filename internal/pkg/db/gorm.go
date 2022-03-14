package db

import (
	"github.com/web3art/g/internal/pkg/dao/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func autoMigrate() error {
	models := []interface{}{}
	models = append(models, model.Tweet{}, model.TweetAuthorAddress{}, model.TweetWaitToClaim{})
	return db.AutoMigrate(models...)
}

func Init(isDev bool, dsn string) error {
	var err error
	if isDev {
		db, err = openSqlite(dsn)
	} else {
		db, err = openMysql(dsn)
	}

	if err != nil {
		return err
	}

	err = autoMigrate()

	return err
}

func DB() *gorm.DB {
	if db == nil {
		panic("db uninitialized")
	}
	return db
}
