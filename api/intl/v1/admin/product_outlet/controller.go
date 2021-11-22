package productoutlet

import (
	"net/http"
	"strconv"

	port "github.com/ikromyalterra/minipos/business/port/product_outlet"
	"github.com/ikromyalterra/minipos/utils/validator"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	prodOutletService port.Service
}

func New(prodOutletService port.Service) *Controller {
	return &Controller{
		prodOutletService,
	}
}

func (controller *Controller) List(c echo.Context) error {
	var outletID uint
	if c.QueryParam("id_outlet") != "" {
		id, err := strconv.ParseUint(c.QueryParam("id_outlet"), 10, 64)
		if err != nil || id == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid id_outlet")
		}
		outletID = uint(id)
	}

	products, err := controller.prodOutletService.List(outletID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	datas := make([]ResponseProductOutletView, 0, len(products))
	var product ResponseProductOutletView
	for i := range products {
		product.ID = products[i].ID
		product.Product = map[string]interface{}{
			"id":    products[i].ProductID,
			"sku":   products[i].ProductSKU,
			"image": products[i].ProductImage,
			"name":  products[i].ProductName,
		}
		product.Outlet = map[string]interface{}{
			"id":   products[i].OutletID,
			"name": products[i].OutletName,
			"merchant": map[string]interface{}{
				"id":   products[i].MerchantID,
				"name": products[i].MerchantName,
			},
		}
		product.Price = products[i].Price

		datas = append(datas, product)
	}

	response := ResponseProductOutletsView{
		ProductOutlet: datas,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Create(c echo.Context) error {
	request := new(RequestProductOutlet)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := validator.GetValidator().Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := new(port.ProductOutlet)
	product.ProductID = request.ProductID
	product.OutletID = request.OutletID
	product.Price = request.Price

	if err := controller.prodOutletService.Create(product); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var response = ResponseProductOutlet{
		ID:        product.ID,
		ProductID: product.ProductID,
		OutletID:  product.OutletID,
		Price:     product.Price,
	}

	return c.JSON(http.StatusCreated, response)
}

func (controller *Controller) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	request := new(RequestProductOutletUpdate)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if request.ProductID > 0 || request.OutletID > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "updating id_product or id_outlet is not allowed")
	}

	product := new(port.ProductOutlet)
	product.ID = uint(id)
	product.Price = request.Price

	if err := controller.prodOutletService.Update(product); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var response = ResponseProductOutlet{
		ID:        product.ID,
		ProductID: product.ProductID,
		OutletID:  product.OutletID,
		Price:     product.Price,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Read(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	product, err := controller.prodOutletService.View(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if product.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "product outlet not found")
	}

	var response ResponseProductOutletView

	response.ID = product.ID
	response.Product = map[string]interface{}{
		"id":    product.ProductID,
		"sku":   product.ProductSKU,
		"image": product.ProductImage,
		"name":  product.ProductName,
	}
	response.Outlet = map[string]interface{}{
		"id":   product.OutletID,
		"name": product.OutletName,
		"merchant": map[string]interface{}{
			"id":   product.MerchantID,
			"name": product.MerchantName,
		},
	}
	response.Price = product.Price

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := controller.prodOutletService.Delete(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}
