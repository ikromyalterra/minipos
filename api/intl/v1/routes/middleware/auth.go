package middleware

import (
	"strings"

	authPort "github.com/ikromyalterra/minipos/business/port/auth"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	authService authPort.Service
}

func NewAuth(authService authPort.Service) *Auth {
	return &Auth{
		authService,
	}
}

func (handler *Auth) CustomParse(tokenString string, c echo.Context) (interface{}, error) {
	return handler.authService.Verify(tokenString)
}

func AuthAPISkipper(c echo.Context) bool {
	paths := []string{
		"/api/v1/auth/login",
	}

	requestURI := c.Request().RequestURI
	for i := range paths {
		if strings.Contains(requestURI, paths[i]) {
			return true
		}
	}
	indexAlwaysAllowed := requestURI == "/" || requestURI == ""

	return indexAlwaysAllowed
}
