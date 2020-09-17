package controller

import (
	"net/http"
	"strconv"

	e "github.com/HETIC-MT-P2021/DB_RAYER_P01/api/entity"
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/api/model"

	"github.com/labstack/echo/v4"
)

// GetCustomer returns a JSON object for one custumer.
func GetCustomer(c echo.Context) error {
	customerNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	res, err := model.GetCustomerByNumber(customerNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, e.SetResponse(http.StatusBadRequest, err.Error(), EmptyValue))
	}

	return c.JSON(http.StatusOK, res)
}
