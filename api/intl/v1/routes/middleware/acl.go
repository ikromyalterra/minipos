package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ikromyalterra/minipos/utils/auth"
	"github.com/labstack/echo/v4"
)

//ACL is method for checking user permisson
func ACL(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			tokenClaims := user.Claims.(*auth.JWTClaims)
			if tokenClaims.Role == permission {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden, "You don't have permission to access this resource")
		}
	}
}
