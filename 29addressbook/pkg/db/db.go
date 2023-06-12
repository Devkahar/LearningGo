package db

import (
	"addressbook/pkg/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB
var DB *gorm.DB

func CreateConnection() {
	dsn := "host=localhost user=postgres password=Dev@123 dbname=stocksdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Contact{}, &model.Address{}, &model.Name{}, &model.Phone{})
	DB = db
}
