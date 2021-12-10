package config

import (
	"awesomeProject/models/DAO"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("failed to connect database")
	}
	database.AutoMigrate(DAO.History{})
	fmt.Println(1)
}
