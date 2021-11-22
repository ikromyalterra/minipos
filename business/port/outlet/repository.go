package outlet

import "github.com/ikromyalterra/minipos/business/model"

// Repository is outbound port
type Repository interface {
	// GetByID get data by ID
	GetByID(ID uint) (model.Outlet, error)

	// View get related data
	View(ID uint) (model.OutletView, error)
}
