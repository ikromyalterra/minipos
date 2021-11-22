package model

import "time"

type (
	UserToken struct {
		TokenID   uint      `json:"id_token" gorm:"primaryKey;column:id_token"`
		UserID    uint      `json:"id_user" gorm:"column:id_user"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
