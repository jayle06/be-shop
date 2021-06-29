package repository

import (
	"ocg.com/project/database"
	"ocg.com/project/model"
	"time"
)

func CreateOder(order *model.Order) *model.Order{
	order.CreatedAt = time.Now().Unix()
	database.DB.Create(&order)
	return order
}

func FindOrderById(id int64) *model.Order {
	order := new(model.Order)
	database.DB.Preload("OrderItem").Where("id = ?", id).Find(&order)
	return order
}
