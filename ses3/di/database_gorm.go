package di

import (
	"github.com/geb/aweproj/ses3/shared/config"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewGorm(conf *config.EnvConfiguration) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(conf.SqlConnectionUri), &gorm.Config{})
	if err != nil {
		panic((err))
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic((err))
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(conf.SqlMaxIdleConnection) * time.Second)
	sqlDb.SetConnMaxLifetime(time.Duration(conf.SqlConnMaxLifetime) * time.Second)
	sqlDb.SetMaxIdleConns(conf.SqlMaxIdleConnection)
	sqlDb.SetMaxOpenConns(2 * conf.SqlMaxIdleConnection)
	return
}

func init() {
	if err := Container.Provide(NewGorm); err != nil {
		panic(errors.Wrap(err, "failed to provide mysql"))
	}
}
