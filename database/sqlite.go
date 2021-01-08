package database

import (
	"restful_api/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	connection, err := gorm.Open(sqlite.Open("./database/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db = connection

	db.AutoMigrate(&entities.User{})

}

// GetSqliteConnectionPool return db
func GetSqliteConnectionPool() *gorm.DB {
	return db
}
