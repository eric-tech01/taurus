package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(args ...interface{}) {
	var err error
	// dbPath := Config.Get("config.dbPath").MustString()
	dbPath := "./base.db"
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	for _, a := range args {
		DB.AutoMigrate(&a)
	}
}
