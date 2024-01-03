package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DB    = gorm.DB
	Model = gorm.Model
)

var (
	// ErrRecordNotFound record not found error.
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

func open(options *Config) (*gorm.DB, error) {
	inner, err := gorm.Open(mysql.Open(options.DSN), &options.gormConfig)
	if err != nil {
		return nil, err
	}

	// inner.(options.Debug)
	// 设置默认连接配置
	db, err := inner.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(options.MaxIdleConns)
	db.SetMaxOpenConns(options.MaxOpenConns)

	if options.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(options.ConnMaxLifetime)
	}

	// 开启 debug
	if options.Debug {
		inner = inner.Debug()
	}

	return inner, err
}
