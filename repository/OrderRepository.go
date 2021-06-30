package repository

import (
	"ocg.com/project/database"
	"ocg.com/project/model"
	"time"
)

func CreateOder(order *model.Order) *model.Order {
	order.CreatedAt = time.Now().Unix()
	database.DB.Create(&order)
	return order
}

func FindOrderById(id int64) *model.Order {
	order := new(model.Order)
	orderItems := new([]model.OrderItem)

	database.DB.Where("order_id = ?", id).Find(&orderItems)
	database.DB.Where("id = ?", id).Find(&order)

	for _, item := range *orderItems {
		if item.OrderId == id {
			order.OrderItems = append(order.OrderItems, item)
		}
	}

	return order
}
