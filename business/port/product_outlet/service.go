package productoutlet

type (
	// ProductOutletOutlet data
	ProductOutlet struct {
		ID           uint
		ProductID    uint
		ProductName  string
		ProductSKU   string
		ProductImage string
		MerchantID   uint
		MerchantName string
		OutletID     uint
		OutletName   string
		Price        uint
	}

	// Service is inbount port
	Service interface {
		// Create insert new data
		Create(productoutlet *ProductOutlet) error

		// GetByID get data by ID
		GetByID(ID uint) (ProductOutlet, error)

		// View get related data
		View(ID uint) (ProductOutlet, error)

		// Update update new data
		Update(productoutlet *ProductOutlet) error

		// Delete delete data
		Delete(ID uint) error

		// List list data
		List(outletID uint) ([]ProductOutlet, error)
	}
)
