package router

import (
	"net/http"
	"regexp"

	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/controller"

	"github.com/labstack/echo/v4"
)

// ParamValidation validate the paramater of the request.
func ParamValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramKey := c.ParamNames()
		paramValue := c.ParamValues()

		r := regexp.MustCompile("^[0-9]+$")

		for k, v := range paramValue {
			if !r.MatchString(v) {
				return c.JSON(http.StatusBadRequest, "param ("+paramKey[k]+") is not a number")
			}
		}

		return next(c)
	}
}

// SetAPIRoutes define all apis routes.
func SetAPIRoutes(e *echo.Echo) {
	// Restricted group
	r := e.Group("/api")

	r.GET("/customer/:id", controller.GetCustomer, ParamValidation)

}
