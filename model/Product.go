package model

type Product struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name"`
	Image       []Image    `json:"image" gorm:"foreignKey:product_id"`
	Price       int64      `json:"price"`
	PromoPrice  int64      `json:"promo_price"`
	IsPromo     bool       `json:"is_promo"`
	Description string     `json:"description"`
	Category    []Category `json:"category" gorm:"many2many:product_category;"`
}