package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Service   string    `json:"service"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
