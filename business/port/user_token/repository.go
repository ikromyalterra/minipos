package user

import "github.com/ikromyalterra/minipos/business/model"

// Repository is outbound port
type Repository interface {
	// Insert insert new data
	Insert(userToken *model.UserToken) error

	// GetByTokenID get data by ID
	GetByTokenID(tokenID uint) (model.UserToken, error)

	// DeleteByUserID delete data
	DeleteByUserID(userID uint) error

	// DeleteByTokenID delete data
	DeleteByTokenID(tokenID uint) error
}
