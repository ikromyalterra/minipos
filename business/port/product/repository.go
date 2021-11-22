package product

import "github.com/ikromyalterra/minipos/business/model"

// Repository is outbound port
type Repository interface {
	// Insert insert new data
	Insert(product *model.Product) error

	// GetByID get data by ID
	GetByID(ID uint) (model.Product, error)

	// GetByMerchantSKU get data by ID
	GetByMerchantSKU(merchantID uint, SKU string) (model.Product, error)

	// View get data by ID and related data
	View(ID uint) (model.ProductView, error)

	// Update update new data
	Update(product model.Product) error

	// Delete delete data
	Delete(ID uint) error

	// List list data
	List(merchantID uint) ([]model.ProductView, error)
}
