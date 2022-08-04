package delivery

import (
	"log"
	"net/http"
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"
	_helper "project3/group3/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase domain.CartUseCase
}

func New(cs domain.CartUseCase) domain.CartHandler {
	return &CartHandler{
		cartUseCase: cs,
	}
}

func (ch *CartHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit := c.QueryParam("limit")
		offset := c.QueryParam("offset")
		limitint, _ := strconv.Atoi(limit)
		offsetint, _ := strconv.Atoi(offset)
		idFromToken, _ := _middleware.ExtractData(c)
		result, err := ch.cartUseCase.GetAllData(limitint, offsetint, idFromToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", FromModelList(result)))
	}
}
func (ch *CartHandler) PostCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		idFromToken, _ := _middleware.ExtractData(c)
		cartReq := InsertFormat{}
		err := c.Bind(&cartReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
		}
		dataCart := domain.Cart{}
		dataCart.Product.ID = cartReq.IdProduct
		dataCart.Stock = cartReq.Stock
		dataCart.Status = "Pending"
		dataCart.UserID = idFromToken
		row, errCreate := ch.cartUseCase.CreateData(dataCart)
		if row == -1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
		}
		if errCreate != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add to cart"))
		}

		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
func (h *CartHandler) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCart, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		cartReq := InsertFormat{}
		err := c.Bind(&cartReq)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
		}
		stock := cartReq.Stock
		row, errUpd := h.cartUseCase.UpdateData(stock, idCart, idFromToken)
		if errUpd != nil {
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row == 0 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}

func (h *CartHandler) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idProd, _ := strconv.Atoi(id)
		idFromToken, _ := _middleware.ExtractData(c)
		row, errDel := h.cartUseCase.DeleteData(idProd, idFromToken)
		if errDel != nil {
			log.Println("cek", errDel)
			return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to delete data user"))
		}
		return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
	}
}
