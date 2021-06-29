package model

type Order struct {
	ID          int64       `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	PhoneNumber string      `json:"phone_number"`
	Address     string      `json:"address"`
	Email       string      `json:"email"`
	Total       float64     `json:"total"`
	OrderItems  []OrderItem `json:"order_items" gorm:"foreignKey:order_id"`
	CreatedAt   int64       `json:"created_at"`
}
