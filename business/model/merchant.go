package model

import "time"

type (
	Merchant struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
