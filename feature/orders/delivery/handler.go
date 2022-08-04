package delivery

import (
	"net/http"
	"project3/group3/domain"
	"strconv"

	_middleware "project3/group3/feature/common"
	_helper "project3/group3/helper"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUseCase domain.OrderUseCase
}

func New(ou domain.OrderUseCase) domain.OrderHandler {
	return &OrderHandler{
		orderUseCase: ou,
	}
}

func (oh *OrderHandler) ConfirmStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idOrder, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, err := oh.orderUseCase.ConfirmData(idOrder, idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (oh *OrderHandler) CancelStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idOrder, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, err := oh.orderUseCase.ConfirmData(idOrder, idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (oh *OrderHandler) GetAllData() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitcnv, _ := strconv.Atoi(limit)
		offsetcnv, _ := strconv.Atoi(offset)
		result, err := oh.orderUseCase.GetAllData(limitcnv, offsetcnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed get all data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(result)))
	}

}
