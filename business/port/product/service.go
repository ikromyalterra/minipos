package product

type (
	// Product data
	Product struct {
		ID           uint
		MerchantID   uint
		MerchantName string
		Image        string
		SKU          string
		Name         string
		Price        uint
	}

	// Service is inbount port
	Service interface {
		// Create insert new data
		Create(product *Product) error

		// GetByID get data by ID
		GetByID(ID uint) (Product, error)

		// View get related data
		View(ID uint) (Product, error)

		// Update update new data
		Update(product *Product) error

		// Delete delete data
		Delete(ID uint) error

		// List list data
		List(merchantID uint) ([]Product, error)
	}
)
