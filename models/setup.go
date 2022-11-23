// models/setup.go

package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {

	// database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// if err != nil {
	//         panic("Failed to connect to database!")
	// }

	// err = database.AutoMigrate(&Book{})
	// if err != nil {
	//         return
	// }

	dsn := "host=localhost user=postgres password=Postgres2022! dbname=postgres port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
