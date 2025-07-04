package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
