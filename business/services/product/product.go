package product

import (
	"errors"

	"github.com/ikromyalterra/minipos/business/model"
	portMerchant "github.com/ikromyalterra/minipos/business/port/merchant"
	port "github.com/ikromyalterra/minipos/business/port/product"
	portProductOutlet "github.com/ikromyalterra/minipos/business/port/product_outlet"
)

type (
	service struct {
		productRepo       port.Repository
		productOutletRepo portProductOutlet.Repository
		merchantRepo      portMerchant.Repository
	}
)

func New(productRepo port.Repository,
	productOutletRepo portProductOutlet.Repository,
	merchantRepo portMerchant.Repository) port.Service {
	return &service{
		productRepo,
		productOutletRepo,
		merchantRepo,
	}
}

func (s *service) Create(product *port.Product) error {
	currentMerchant, err := s.merchantRepo.GetByID(product.MerchantID)
	if err != nil || currentMerchant.ID == 0 {
		return errors.New("id_merchant not found")
	}

	existingProduct, err := s.productRepo.GetByMerchantSKU(product.MerchantID, product.SKU)
	if err != nil {
		return err
	}
	if existingProduct.ID > 0 {
		return errors.New("sku already used by other product")
	}

	data := new(model.Product)
	data.SKU = product.SKU
	data.Image = product.Image
	data.MerchantID = product.MerchantID
	data.Name = product.Name
	data.Price = product.Price

	if err := s.productRepo.Insert(data); err != nil {
		return err
	}
	product.ID = data.ID

	return nil
}

func (s *service) GetByID(ID uint) (port.Product, error) {
	var data port.Product

	product, err := s.productRepo.GetByID(ID)
	if product.ID > 0 {
		data.ID = product.ID
		data.SKU = product.SKU
		data.Image = product.Image
		data.MerchantID = product.MerchantID
		data.Name = product.Name
		data.Price = product.Price
	}

	return data, err
}

func (s *service) View(ID uint) (port.Product, error) {
	var data port.Product

	product, err := s.productRepo.View(ID)
	if product.ID > 0 {
		data.ID = product.ID
		data.SKU = product.SKU
		data.Image = product.Image
		data.MerchantID = product.MerchantID
		data.MerchantName = product.MerchantName
		data.Name = product.Name
		data.Price = product.Price
	}

	return data, err
}

func (s *service) Update(product *port.Product) error {
	currentProduct, err := s.productRepo.GetByID(product.ID)
	if err != nil {
		return err
	}
	if currentProduct.ID == 0 {
		return errors.New("id not found")
	}

	if currentProduct.MerchantID != product.MerchantID {
		productOutlet, err := s.productOutletRepo.GetByProductID(product.ID)
		if productOutlet.ID > 0 || err != nil {
			return errors.New("Can't change id_merchant, this product already defined in outlet")
		}
	}
	if currentProduct.SKU != product.SKU {
		existingProduct, err := s.productRepo.GetByMerchantSKU(product.MerchantID, product.SKU)
		if existingProduct.ID > 0 || err != nil {
			return errors.New("SKU already used by other product")
		}
	}

	var data model.Product
	data.ID = product.ID
	data.SKU = product.SKU
	data.Image = product.Image
	data.MerchantID = product.MerchantID
	data.Name = product.Name
	data.Price = product.Price
	data.CreatedAt = currentProduct.CreatedAt

	return s.productRepo.Update(data)
}

func (s *service) Delete(ID uint) error {
	return s.productRepo.Delete(ID)
}

func (s *service) List(merchantID uint) ([]port.Product, error) {
	datas := make([]port.Product, 0)

	products, err := s.productRepo.List(merchantID)
	if err != nil {
		return datas, err
	}

	var product port.Product
	for i := range products {
		product.ID = products[i].ID
		product.SKU = products[i].SKU
		product.Image = products[i].Image
		product.MerchantID = products[i].MerchantID
		product.MerchantName = products[i].MerchantName
		product.Name = products[i].Name
		product.Price = products[i].Price

		datas = append(datas, product)
	}
	return datas, nil
}
