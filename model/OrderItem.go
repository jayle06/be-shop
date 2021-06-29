package model

type OrderItem struct {
	ID          int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderId     int64   `json:"order_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

