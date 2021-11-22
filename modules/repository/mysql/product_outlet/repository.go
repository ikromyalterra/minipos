package productoutlet

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
		db.Table("product_outlet"),
	}
}

func (repo *Repository) Insert(productOutlet *model.ProductOutlet) error {
	return repo.Create(productOutlet).Error
}

func (repo *Repository) GetByID(ID uint) (model.ProductOutlet, error) {
	var productOutlet model.ProductOutlet
	result := repo.Where("id = ?", ID).Find(&productOutlet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return productOutlet, result.Error
}

func (repo *Repository) GetByProductID(productID uint) (model.ProductOutlet, error) {
	var productOutlet model.ProductOutlet
	result := repo.Where("id_product = ?", productID).Limit(1).Find(&productOutlet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return productOutlet, result.Error
}

func (repo *Repository) View(ID uint) (model.ProductOutletView, error) {
	var productOutlet model.ProductOutletView
	result := repo.Select("product_outlet.*, product.sku as sku, product.image as image,product.name as product, merchant.id as id_merchant, merchant.name as merchant, outlet.name as outlet").
		Joins("join product on product_outlet.id_product = product.id").
		Joins("join outlet on product_outlet.id_outlet = outlet.id").
		Joins("join merchant on product.id_merchant = merchant.id").
		Where("product_outlet.id = ?", ID).
		Find(&productOutlet)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result.Error = nil
	}
	return productOutlet, result.Error
}

func (repo *Repository) Update(productOutlet model.ProductOutlet) error {
	mapProductOutlet, err := helper.StructToMap(productOutlet)
	if err != nil {
		return err
	}

	return repo.Model(&productOutlet).Updates(mapProductOutlet).Error
}

func (repo *Repository) Delete(ID uint) error {
	result := repo.Where("id = ?", ID).Delete(&model.ProductOutlet{})
	if result.RowsAffected == 0 {
		return errors.New("product outlet not found")
	}
	return result.Error
}

func (repo *Repository) List(outletID uint) ([]model.ProductOutletView, error) {
	var productOutlets []model.ProductOutletView
	q := repo.Select("product_outlet.*, product.sku as sku, product.image as image, product.name as product, merchant.id as id_merchant, merchant.name as merchant, outlet.name as outlet").
		Joins("join product on product_outlet.id_product = product.id").
		Joins("join outlet on product_outlet.id_outlet = outlet.id").
		Joins("join merchant on product.id_merchant = merchant.id")
	if outletID > 0 {
		q = q.Where("product_outlet.id_outlet = ?", outletID)
	}
	q.Find(&productOutlets)

	return productOutlets, nil
}
