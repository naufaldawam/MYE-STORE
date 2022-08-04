package usecase

import (
	"errors"
	"project3/group3/domain"
)

type orderUseCase struct {
	orderData domain.OrderData
}

func New(od domain.OrderData) domain.OrderUseCase {
	return &orderUseCase{
		orderData: od,
	}
}

func (uo *orderUseCase) CreateData(data domain.Cart) (row int, err error) {
	if data.Stock == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	return row, nil
}

func (uo *orderUseCase) ConfirmData(idOrder, idFromToken int) (row int, err error) {
	row, err = uo.orderData.ConfirmOrder(idOrder, idFromToken)
	return row, err
}

func (uo *orderUseCase) CancelData(idOrder, idFromToken int) (row int, err error) {
	row, err = uo.orderData.ConfirmOrder(idOrder, idFromToken)
	return row, err
}

func (uo *orderUseCase) GetAllData(limit, offset int) (data []domain.Order, err error) {
	data, err = uo.orderData.SelectData(limit, offset)
	for k, v := range data {
		data[k].TotalPrice = v.Stock * v.Product.Price
	}
	return data, err
}
