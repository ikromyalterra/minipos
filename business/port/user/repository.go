package user

import "github.com/ikromyalterra/minipos/business/model"

// Repository is outbound port
type Repository interface {
	// Insert insert new data
	Insert(user *model.User) error

	// Read get data by ID
	GetByID(ID uint) (model.User, error)

	// View get data related by IDs
	View(ID uint) (model.UserView, error)

	// GetByEmail find with email
	GetByEmail(email string) (model.User, error)

	// Update update new data
	Update(user model.User) error

	// Delete delete data
	Delete(ID uint) error

	// List list data
	List() ([]model.UserView, error)
}
