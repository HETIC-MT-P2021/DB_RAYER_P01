package controller

import (
	"net/http"
	"strconv"

	"github.com/HETIC-MT-P2021/DB_RAYER_P01/helper"
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/model"

	"github.com/labstack/echo/v4"
)

// GetOrderProduct returns a JSON object for one order.
func GetOrderProduct(c echo.Context) error {
	orderNumber, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, err.Error(), helper.EmptyValue))
	}

	res, err := model.GetProductByOrderNumber(orderNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, err.Error(), helper.EmptyValue))
	}

	return c.JSON(http.StatusOK, res)
}
