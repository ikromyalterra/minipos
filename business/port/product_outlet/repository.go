package productoutlet

import "github.com/ikromyalterra/minipos/business/model"

// Repository is outbound port
type Repository interface {
	// Insert insert new data
	Insert(productOutlet *model.ProductOutlet) error

	// GetByID get data by ID
	GetByID(ID uint) (model.ProductOutlet, error)

	// GetByProductID check product already in outlet
	GetByProductID(productID uint) (model.ProductOutlet, error)

	// View get data by ID
	View(ID uint) (model.ProductOutletView, error)

	// Update update new data
	Update(productOutlet model.ProductOutlet) error

	// Delete delete data
	Delete(ID uint) error

	// List list data
	List(outletID uint) ([]model.ProductOutletView, error)
}
