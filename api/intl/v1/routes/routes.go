package routes

import (
	"github.com/ikromyalterra/minipos/utils/config"

	m "github.com/ikromyalterra/minipos/api/intl/v1/routes/middleware"

	merchantRepository "github.com/ikromyalterra/minipos/modules/repository/mysql/merchant"
	outletRepository "github.com/ikromyalterra/minipos/modules/repository/mysql/outlet"
	productRepository "github.com/ikromyalterra/minipos/modules/repository/mysql/product"
	productOutletRepository "github.com/ikromyalterra/minipos/modules/repository/mysql/product_outlet"
	userRepository "github.com/ikromyalterra/minipos/modules/repository/mysql/user"
	userTokenRepository "github.com/ikromyalterra/minipos/modules/repository/mysql/user_token"

	authService "github.com/ikromyalterra/minipos/business/services/auth"
	productService "github.com/ikromyalterra/minipos/business/services/product"
	productOutletService "github.com/ikromyalterra/minipos/business/services/product_outlet"
	userService "github.com/ikromyalterra/minipos/business/services/user"

	adminProductCtrl "github.com/ikromyalterra/minipos/api/intl/v1/admin/product"
	adminProductOutletCtrl "github.com/ikromyalterra/minipos/api/intl/v1/admin/product_outlet"
	adminUserCtrl "github.com/ikromyalterra/minipos/api/intl/v1/admin/user"

	authCtrl "github.com/ikromyalterra/minipos/api/intl/v1/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func acl(permission string) echo.MiddlewareFunc {
	return m.ACL(permission)
}

func API(e *echo.Echo) {
	// Instance DB
	db := config.DB

	userRepo := userRepository.New(db)
	userTokenRepo := userTokenRepository.New(db)

	authServ := authService.New(userTokenRepo, userRepo)

	// set jwt
	customMiddleware := m.NewAuth(authServ)
	JWTCustomConfig := middleware.JWTConfig{
		Skipper:        m.AuthAPISkipper,
		ParseTokenFunc: customMiddleware.CustomParse,
	}

	// instance repo
	merchantRepo := merchantRepository.New(db)
	outletRepo := outletRepository.New(db)
	productRepo := productRepository.New(db)
	productOutletRepo := productOutletRepository.New(db)

	// instance service
	productOutletServ := productOutletService.New(productOutletRepo, productRepo, outletRepo)
	productServ := productService.New(productRepo, productOutletRepo, merchantRepo)
	userServ := userService.New(userRepo, merchantRepo, outletRepo, userTokenRepo)

	e.Use(middleware.JWTWithConfig(JWTCustomConfig))

	// auth route
	authHandler := authCtrl.New(authServ)
	authRoute := e.Group("/api/v1/auth")
	authRoute.POST("/login", authHandler.Login)
	authRoute.POST("/logout", authHandler.Logout)

	// admin routes
	// user
	adminUserHandler := adminUserCtrl.New(userServ)
	adminUserRoute := e.Group("/api/v1/admin/user", acl("admin"))
	adminUserRoute.POST("", adminUserHandler.Create)
	adminUserRoute.GET("/:id", adminUserHandler.Read)
	adminUserRoute.PUT("/:id", adminUserHandler.Update)
	adminUserRoute.DELETE("/:id", adminUserHandler.Delete)
	adminUserRoute.GET("", adminUserHandler.List)
	// product outlet
	adminProductOutletHandler := adminProductOutletCtrl.New(productOutletServ)
	adminProductOutletRoute := e.Group("/api/v1/admin/product/outlet", acl("admin"))
	adminProductOutletRoute.POST("", adminProductOutletHandler.Create)
	adminProductOutletRoute.GET("/:id", adminProductOutletHandler.Read)
	adminProductOutletRoute.PUT("/:id", adminProductOutletHandler.Update)
	adminProductOutletRoute.DELETE("/:id", adminProductOutletHandler.Delete)
	adminProductOutletRoute.GET("", adminProductOutletHandler.List)
	// product
	adminProductHandler := adminProductCtrl.New(productServ)
	adminProductRoute := e.Group("/api/v1/admin/product", acl("admin"))
	adminProductRoute.POST("", adminProductHandler.Create)
	adminProductRoute.GET("/:id", adminProductHandler.Read)
	adminProductRoute.PUT("/:id", adminProductHandler.Update)
	adminProductRoute.DELETE("/:id", adminProductHandler.Delete)
	adminProductRoute.GET("", adminProductHandler.List)
}
