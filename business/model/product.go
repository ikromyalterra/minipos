package model

import "time"

type (
	Product struct {
		ID         uint      `json:"id" gorm:"primaryKey"`
		MerchantID uint      `json:"id_merchant" gorm:"column:id_merchant"`
		SKU        string    `json:"sku"`
		Image      string    `json:"image" gorm:"default:''"`
		Name       string    `json:"name"`
		Price      uint      `json:"price" gorm:"default:0"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	ProductView struct {
		ID           uint      `json:"id" gorm:"primaryKey"`
		MerchantID   uint      `json:"id_merchant" gorm:"column:id_merchant"`
		MerchantName string    `json:"merchant" gorm:"column:merchant;default:null"`
		SKU          string    `json:"sku"`
		Image        string    `json:"image" gorm:"default:''"`
		Name         string    `json:"name"`
		Price        uint      `json:"price" gorm:"default:0"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
