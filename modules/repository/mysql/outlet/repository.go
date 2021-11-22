package outlet

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
		db.Table("outlet"),
	}
}

func (repo *Repository) GetByID(ID uint) (model.Outlet, error) {
	var outlet model.Outlet
	result := repo.Where("id = ?", ID).Find(&outlet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return outlet, result.Error
}

func (repo *Repository) View(ID uint) (model.OutletView, error) {
	var outlet model.OutletView
	result := repo.Select("outlet.*, merchant.name as merchant").
		Joins("join merchant on outlet.id_merchant = merchant.id").
		Where("outlet.id = ?", ID).
		Find(&outlet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return outlet, result.Error
}
