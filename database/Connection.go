package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ocg.com/project/model"
)

var DB *gorm.DB

func Connect(){
	database, err := gorm.Open(mysql.Open("root:abc123@tcp(127.0.0.1:3305)/project?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database!")
	}

	DB = database

	database.AutoMigrate(&model.Product{}, &model.Image{}, &model.Category{}, &model.User{}, &model.Order{}, &model.OrderItem{})
}