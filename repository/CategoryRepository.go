package repository

import (
	"ocg.com/project/database"
	"ocg.com/project/model"
)

func FindAllCategories() *[]model.Category {
	categories := new([]model.Category)
	database.DB.Find(&categories)
	return categories
}


