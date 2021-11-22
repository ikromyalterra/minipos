package user

import (
	"net/http"
	"strconv"

	port "github.com/ikromyalterra/minipos/business/port/user"
	"github.com/ikromyalterra/minipos/utils/validator"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	userService port.Service
}

func New(userService port.Service) *Controller {
	return &Controller{
		userService,
	}
}

func (controller *Controller) List(c echo.Context) error {
	users, err := controller.userService.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	datas := make([]*ResponseUserView, 0, len(users))
	for i := range users {
		user := populateResponseUserView(&users[i])
		datas = append(datas, user)
	}

	response := ResponseUsersView{
		User: datas,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) Create(c echo.Context) error {
	request := new(RequestUser)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := new(port.User)
	user.Email = request.Email
	user.Password = request.Password
	user.Role = request.Role
	user.MerchantID = request.MerchantID
	user.OutletID = request.OutletID

	if err := controller.userService.Create(user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusCreated, populateResponseUser(user))
}

func (controller *Controller) Update(c echo.Context) error {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	request := new(RequestUserUpdate)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := validator.GetValidator().Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := new(port.User)
	user.ID = uint(uid)
	user.Email = request.Email
	user.Password = request.Password
	user.Role = request.Role
	user.MerchantID = request.MerchantID
	user.OutletID = request.OutletID

	if err := controller.userService.Update(user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, populateResponseUser(user))
}

func (controller *Controller) Read(c echo.Context) error {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil || uid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	user, err := controller.userService.View(uint(uid))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, populateResponseUserView(&user))
}

func (controller *Controller) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := controller.userService.Delete(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "")
}

func populateResponseUserView(user *port.User) *ResponseUserView {
	response := new(ResponseUserView)
	var responseUserMerchant map[string]interface{}
	var responseUserOutlet map[string]interface{}

	if user.MerchantID > 0 {
		responseUserMerchant = map[string]interface{}{
			"id":   user.MerchantID,
			"name": user.MerchantName,
		}
	}
	if user.OutletID > 0 {
		responseUserOutlet = map[string]interface{}{
			"id":   user.OutletID,
			"name": user.OutletName,
		}
	}
	response.ID = user.ID
	response.Email = user.Email
	response.Role = user.Role
	response.Merchant = responseUserMerchant
	response.Outlet = responseUserOutlet

	return response
}

func populateResponseUser(user *port.User) *ResponseUser {
	response := new(ResponseUser)

	response.ID = user.ID
	response.Email = user.Email
	response.Role = user.Role
	response.MerchantID = user.MerchantID
	response.OutletID = user.OutletID
	if user.MerchantID == 0 {
		response.MerchantID = nil
	}
	if user.OutletID == 0 {
		response.OutletID = nil
	}

	return response
}
