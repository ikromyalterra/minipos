package productoutlet

import (
	"errors"

	"github.com/ikromyalterra/minipos/business/model"
	portOutlet "github.com/ikromyalterra/minipos/business/port/outlet"
	portProduct "github.com/ikromyalterra/minipos/business/port/product"
	port "github.com/ikromyalterra/minipos/business/port/product_outlet"
)

type (
	service struct {
		productOutletRepo port.Repository
		productRepo       portProduct.Repository
		outletRepo        portOutlet.Repository
	}
)

func New(productOutletRepo port.Repository,
	productRepo portProduct.Repository,
	outletRepo portOutlet.Repository) port.Service {
	return &service{
		productOutletRepo,
		productRepo,
		outletRepo,
	}
}

func (s *service) Create(productOutlet *port.ProductOutlet) error {
	product, err := s.productRepo.GetByID(productOutlet.ProductID)
	if err != nil || product.ID == 0 {
		return errors.New("id_product not found")
	}
	outlet, err := s.outletRepo.GetByID(productOutlet.OutletID)
	if err != nil || outlet.ID == 0 {
		return errors.New("id_outlet not found")
	}
	if product.MerchantID != outlet.MerchantID {
		return errors.New("invalid data between merchant and outlet")
	}

	data := new(model.ProductOutlet)
	data.ProductID = productOutlet.ProductID
	data.OutletID = productOutlet.OutletID
	data.Price = productOutlet.Price
	if err := s.productOutletRepo.Insert(data); err != nil {
		return err
	}
	productOutlet.ID = data.ID

	return nil
}

func (s *service) GetByID(ID uint) (port.ProductOutlet, error) {
	var data port.ProductOutlet

	productOutlet, err := s.productOutletRepo.GetByID(ID)
	if productOutlet.ID > 0 {
		data.ID = productOutlet.ID
		data.ProductID = productOutlet.ProductID
		data.OutletID = productOutlet.OutletID
		data.Price = productOutlet.Price
	}

	return data, err
}

func (s *service) View(ID uint) (port.ProductOutlet, error) {
	var data port.ProductOutlet

	productOutlet, err := s.productOutletRepo.View(ID)
	if productOutlet.ID > 0 {
		data.ID = productOutlet.ID
		data.ProductID = productOutlet.ProductID
		data.ProductName = productOutlet.ProductName
		data.ProductSKU = productOutlet.ProductSKU
		data.ProductImage = productOutlet.ProductImage
		data.MerchantID = productOutlet.MerchantID
		data.MerchantName = productOutlet.MerchantName
		data.OutletID = productOutlet.OutletID
		data.OutletName = productOutlet.OutletName
		data.Price = productOutlet.Price
	}

	return data, err
}

func (s *service) Update(productOutlet *port.ProductOutlet) error {
	currentProductOutlet, err := s.productOutletRepo.GetByID(productOutlet.ID)
	if err != nil {
		return err
	}
	if currentProductOutlet.ID == 0 {
		return errors.New("id not found")
	}

	productOutlet.ProductID = currentProductOutlet.ProductID
	productOutlet.OutletID = currentProductOutlet.OutletID

	var data model.ProductOutlet
	data.ID = productOutlet.ID
	data.Price = productOutlet.Price
	data.ProductID = productOutlet.ProductID
	data.OutletID = productOutlet.OutletID
	data.CreatedAt = currentProductOutlet.CreatedAt

	return s.productOutletRepo.Update(data)
}

func (s *service) Delete(ID uint) error {
	return s.productOutletRepo.Delete(ID)
}

func (s *service) List(outletID uint) ([]port.ProductOutlet, error) {
	datas := make([]port.ProductOutlet, 0)

	productOutlets, err := s.productOutletRepo.List(outletID)
	if err != nil {
		return datas, err
	}

	var productOutlet port.ProductOutlet
	for i := range productOutlets {
		productOutlet.ID = productOutlets[i].ID
		productOutlet.ProductID = productOutlets[i].ProductID
		productOutlet.ProductName = productOutlets[i].ProductName
		productOutlet.ProductSKU = productOutlets[i].ProductSKU
		productOutlet.ProductImage = productOutlets[i].ProductImage
		productOutlet.MerchantID = productOutlets[i].MerchantID
		productOutlet.MerchantName = productOutlets[i].MerchantName
		productOutlet.OutletID = productOutlets[i].OutletID
		productOutlet.OutletName = productOutlets[i].OutletName
		productOutlet.Price = productOutlets[i].Price

		datas = append(datas, productOutlet)
	}

	return datas, nil
}
