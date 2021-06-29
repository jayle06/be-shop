package model

type Image struct {
	Id        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	ImageUrl  string `json:"image_url"`
	ProductId int64  `json:"product_id"`
}
