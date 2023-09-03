package db

import (
	"jwt/auth/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func Connection() {
	dsn := "host=localhost user=postgres password=Dev@123 dbname=jwtauth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	DBInstance = db
}
