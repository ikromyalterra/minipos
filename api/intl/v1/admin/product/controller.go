package product

import (
	"net/http"
	"strconv"

	port "github.com/ikromyalterra/minipos/business/port/product"
	"github.com/ikromyalterra/minipos/utils/validator"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	productService port.Service
}

func New(productService port.Service) *Controller {
	return &Controller{
		productService,
	}
}

func (controller *Controller) List(c echo.Context) error {
	var merchantID uint
	if c.QueryParam("id_merchant") != "" {
		id, err := strconv.ParseUint(c.QueryParam("id_merchant"), 10, 64)
		if err != nil || id == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id_merchant")
		}
		merchantID = uint(id)
	}
	products, err := controller.productService.List(merchantID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	datas := make([]ResponseProductView, 0, len(products))
	var product ResponseProductView
	for i := range products {
		product.ID = products[i].ID
		product.SKU = products[i].SKU
		product.Image = products[i].Image
		product.Name = products[i].Name
		product.Merchant = map[string]interface{}{
			"id":   products[i].MerchantID,
			"name": products[i].MerchantName,
		}
		product.Price = products[i].Price

		datas = append(datas, product)
	}

	response := ResponseProductsView{
		Product: datas,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Create(c echo.Context) error {
	request := new(RequestProduct)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := new(port.Product)
	product.SKU = request.SKU
	product.Image = request.Image
	product.Name = request.Name
	product.MerchantID = request.MerchantID
	product.Price = request.Price

	if err := controller.productService.Create(product); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var response = ResponseProduct{
		ID:         product.ID,
		SKU:        product.SKU,
		Image:      product.Image,
		Name:       product.Name,
		MerchantID: product.MerchantID,
		Price:      product.Price,
	}

	return c.JSON(http.StatusCreated, response)
}

func (controller *Controller) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	request := new(RequestProduct)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := new(port.Product)
	product.ID = uint(id)
	product.SKU = request.SKU
	product.Image = request.Image
	product.Name = request.Name
	product.Price = request.Price
	product.MerchantID = request.MerchantID

	if err := controller.productService.Update(product); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var response = ResponseProduct{
		ID:         product.ID,
		SKU:        product.SKU,
		Image:      product.Image,
		Name:       product.Name,
		MerchantID: product.MerchantID,
		Price:      product.Price,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Read(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	product, err := controller.productService.View(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if product.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "product not found")
	}

	var response = ResponseProductView{
		ID:    product.ID,
		SKU:   product.SKU,
		Image: product.Image,
		Name:  product.Name,
		Merchant: map[string]interface{}{
			"id":   product.MerchantID,
			"name": product.MerchantName,
		},
		Price: product.Price,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := controller.productService.Delete(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
