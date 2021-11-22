package model

import "time"

type (
	ProductOutlet struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		ProductID uint      `json:"id_product" gorm:"column:id_product"`
		OutletID  uint      `json:"id_outlet" gorm:"column:id_outlet"`
		Price     uint      `json:"price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ProductOutletView struct {
		ID           uint      `json:"id" gorm:"primaryKey"`
		ProductID    uint      `json:"id_product" gorm:"column:id_product"`
		ProductName  string    `json:"product" gorm:"column:product"`
		ProductSKU   string    `json:"sku" gorm:"column:sku"`
		ProductImage string    `json:"image" gorm:"column:image"`
		MerchantID   uint      `json:"id_merchant" gorm:"column:id_merchant"`
		MerchantName string    `json:"merchant" gorm:"column:merchant"`
		OutletID     uint      `json:"id_outlet" gorm:"column:id_outlet"`
		OutletName   string    `json:"outlet" gorm:"column:outlet"`
		Price        uint      `json:"price"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
