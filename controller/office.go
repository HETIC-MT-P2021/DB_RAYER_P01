package controller

import (
	"net/http"

	"github.com/HETIC-MT-P2021/DB_RAYER_P01/helper"
	"github.com/HETIC-MT-P2021/DB_RAYER_P01/model"

	"github.com/labstack/echo/v4"
)

// GetAllOffice returns a JSON list of office.
func GetAllOffice(c echo.Context) error {
	res, err := model.GetAllOffice()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, err.Error(), helper.EmptyValue))
	}

	if len(res) == 0 {
		return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "Employee is empty", helper.EmptyValue))
	}

	return c.JSON(http.StatusOK, res)
}
