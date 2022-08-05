package usecase

import (
	"errors"
	"project3/group3/domain"
)

type cartUseCase struct {
	cartData domain.ChartData
}

func New(pd domain.ChartData) domain.CartUseCase {
	return &cartUseCase{
		cartData: pd,
	}
}

func (uc *cartUseCase) GetAllData(limit, offset, idFromToken int) (data []domain.Cart, err error) {
	data, err = uc.cartData.SelectData(limit, offset, idFromToken)
	for k, v := range data {
		data[k].TotalPrice = v.Stock * v.Product.Price
	}
	return data, err
}

func (uc *cartUseCase) CreateData(data domain.Cart) (row int, err error) {
	if data.Stock == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	isExist, idCart, stock, _ := uc.cartData.CheckCart(data.Product.ID, data.UserID)
	if isExist {
		row, err = uc.cartData.UpdateDataDB(stock+1, idCart, data.UserID)
	} else {
		row, err = uc.cartData.InsertData(data)
	}

	return row, err
}

func (uc *cartUseCase) UpdateData(stock, idCart, idFromToken int) (row int, err error) {
	row, err = uc.cartData.UpdateDataDB(stock, idCart, idFromToken)
	return row, err
}

func (uc *cartUseCase) DeleteData(idCart, idFromToken int) (row int, err error) {
	row, err = uc.cartData.DeleteDataDB(idCart, idFromToken)
	return row, err
}
