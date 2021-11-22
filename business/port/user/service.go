package user

type (
	// User data
	User struct {
		ID           uint
		Email        string
		Password     string
		Role         string
		MerchantID   uint
		MerchantName string
		OutletID     uint
		OutletName   string
	}

	// Service is inbount port
	Service interface {
		// Create insert new data
		Create(user *User) error

		// View get data by ID with related data
		View(ID uint) (User, error)

		// GetByID get data by ID
		GetByID(ID uint) (User, error)

		// Update update new data
		Update(user *User) error

		// Delete delete data
		Delete(ID uint) error

		// List list data
		List() ([]User, error)
	}
)
