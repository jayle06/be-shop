package repository

import (
	"gorm.io/gorm/clause"
	"ocg.com/project/database"
	"ocg.com/project/model"
)

func CreateProduct(product *model.Product) *model.Product {
	database.DB.Create(&product)
	return product
}

func FindAllProducts() *[]model.Product {
	products := new([]model.Product)
	database.DB.Preload("Image").Preload("Category").Find(&products)
	return products
}

func FindProductById(id int64) *model.Product {
	product := new(model.Product)
	database.DB.Preload("Image").Preload("Category").Where("id = ?", id).Find(&product)
	return product
}

func DeleteProductById(id int64) {
	product := new(model.Product)
	database.DB.Where("id = ?", id).Select(clause.Associations).Delete(&product)
}

func FindProductByCategory(categoryName string) *[]model.Product {
	products := new([]model.Product)
	database.DB.Joins("JOIN product_category on " +
		"product_category.product_id=products.id").Joins("JOIN categories on " +
			"categories.id=product_category.category_id").Where("categories.name = " +
				"?", categoryName).Preload("Image").Preload("Category").Find(&products)

	return products
}

