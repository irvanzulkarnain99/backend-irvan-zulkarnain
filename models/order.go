package models

import "time"

type Order struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	ProductID     uint      `json:"product_id"`
	MerchantID    uint      `json:"merchant_id"`
	Quantity      int       `json:"quantity"`
	TotalPrice    int       `json:"total_price"`
	ShippingPrice int       `json:"shipping_price"`
	Discount      int       `json:"discount"`
	CreatedAt     time.Time `json:"created_at"`
}
