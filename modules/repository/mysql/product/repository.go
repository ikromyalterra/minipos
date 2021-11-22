package product

import (
	"errors"

	"github.com/ikromyalterra/minipos/business/model"
	"github.com/ikromyalterra/minipos/utils/helper"
	"github.com/jinzhu/gorm"
)

type (
	Repository struct {
		*gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		db.Table("product"),
	}
}

func (repo *Repository) Insert(product *model.Product) error {
	return repo.Create(product).Error
}

func (repo *Repository) GetByID(ID uint) (model.Product, error) {
	var product model.Product
	result := repo.Where("id = ?", ID).Find(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return product, result.Error
}

func (repo *Repository) GetByMerchantSKU(merchantID uint, SKU string) (model.Product, error) {
	var product model.Product
	result := repo.Where("id_merchant = ? AND sku = ?", merchantID, SKU).Find(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return product, result.Error
}

func (repo *Repository) View(ID uint) (model.ProductView, error) {
	var product model.ProductView
	result := repo.Select("product.*, merchant.name as merchant").
		Joins("join merchant on product.id_merchant = merchant.id").
		Where("product.id = ?", ID).
		Find(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return product, result.Error
}

func (repo *Repository) Update(product model.Product) error {
	mapProduct, err := helper.StructToMap(product)
	if err != nil {
		return err
	}

	return repo.Model(&product).Updates(mapProduct).Error
}

func (repo *Repository) Delete(ID uint) error {
	result := repo.Where("id = ?", ID).Delete(&model.Product{})
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return result.Error
}

func (repo *Repository) List(merchantID uint) ([]model.ProductView, error) {
	var products []model.ProductView

	q := repo.Select("product.*, merchant.name as merchant").
		Joins("join merchant on product.id_merchant = merchant.id")

	if merchantID > 0 {
		q = q.Where("product.id_merchant = ?", merchantID)
	}
	q.Find(&products)

	return products, nil
}
