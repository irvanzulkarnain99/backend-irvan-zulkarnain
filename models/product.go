package models

type Product struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	MerchantID uint   `json:"merchant_id"`
}
