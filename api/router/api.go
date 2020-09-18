package router

import (
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/controller"
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/helper"

	"github.com/labstack/echo/v4"
)

// SetAPIRoutes define all apis routes.
func SetAPIRoutes(e *echo.Echo) {
	// Restricted group
	r := e.Group("/api")

	r.GET("/customer/:id", controller.GetCustomer, helper.ParamValidation)
	r.GET("/order/:id/product", controller.GetOrderProduct, helper.ParamValidation)
	r.GET("/employee", controller.GetAllEmployee)

}
