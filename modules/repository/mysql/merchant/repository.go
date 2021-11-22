package merchant

import (
	"errors"

	"github.com/ikromyalterra/minipos/business/model"
	"github.com/jinzhu/gorm"
)

type (
	Repository struct {
		*gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		db.Table("merchant"),
	}
}

func (repo *Repository) GetByID(ID uint) (model.Merchant, error) {
	var merchant model.Merchant
	result := repo.Where("id = ?", ID).Find(&merchant)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return merchant, result.Error
}
