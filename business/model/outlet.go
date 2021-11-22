package model

import "time"

type (
	Outlet struct {
		ID         uint      `json:"id" gorm:"primaryKey"`
		MerchantID uint      `json:"id_merchant" gorm:"column:id_merchant"`
		Name       string    `json:"name"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	OutletView struct {
		ID           uint      `json:"id" gorm:"primaryKey"`
		MerchantID   uint      `json:"id_merchant" gorm:"column:id_merchant"`
		MerchantName string    `json:"merchant" gorm:"column:merchant"`
		Name         string    `json:"name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
