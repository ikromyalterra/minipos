package merchant

import "github.com/ikromyalterra/minipos/business/model"

// Repository is outbound port
type Repository interface {

	// Read get data by ID
	GetByID(ID uint) (model.Merchant, error)
}
