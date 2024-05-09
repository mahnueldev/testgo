package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}
