package model

import "time"

type (
	User struct {
		ID         uint      `json:"id" gorm:"primaryKey"`
		Email      string    `json:"email"`
		Password   string    `json:"password,omitempty"`
		Role       string    `json:"role"`
		MerchantID uint      `json:"id_merchant" gorm:"column:id_merchant;default:null"`
		OutletID   uint      `json:"id_outlet" gorm:"column:id_outlet;default:null"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	UserView struct {
		ID           uint      `json:"id" gorm:"primaryKey"`
		Email        string    `json:"email"`
		Password     string    `json:"password,omitempty"`
		Role         string    `json:"role"`
		MerchantID   uint      `json:"id_merchant" gorm:"column:id_merchant;default:null"`
		MerchantName string    `json:"merchant" gorm:"column:merchant;default:null"`
		OutletID     uint      `json:"id_outlet" gorm:"column:id_outlet;default:null"`
		OutletName   string    `json:"outlet" gorm:"column:outlet;default:null"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
