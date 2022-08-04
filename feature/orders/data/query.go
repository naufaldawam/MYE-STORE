package data

import (
	"errors"
	"log"
	"project3/group3/domain"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderData {
	return &orderData{
		db: db,
	}
}

func (od *orderData) InsertData(data domain.Order) (row int, err error) {
	order := FromDomain(data)
	result := od.db.Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to create data order")
	}
	return int(result.RowsAffected), nil
}

func (od *orderData) UpdateDataDB(stock, idOrder, idFromToken int) (row int, err error) {
	dataOrder := Order{}
	idCheck := od.db.First(&dataOrder, idOrder)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataOrder.UserID != idFromToken {
		log.Println("check", dataOrder.UserID)
		return -1, errors.New("you don't have access")
	}
	result := od.db.Model(&Order{}).Where("id = ?", idOrder).Update("stock", stock)

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}
	return int(result.RowsAffected), nil
}

func (od *orderData) ConfirmOrder(idOrder, idFromToken int) (row int, err error) {
	dataOrder := Order{}
	idCheck := od.db.First(&dataOrder, idOrder)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataOrder.UserID != idFromToken {
		return -1, errors.New("you don't have access")
	}
	result := od.db.Model(&Order{}).Where("id = ?", idOrder).Update("status", "Confirmed")

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update order status")
	}
	return int(result.RowsAffected), nil
}

func (od *orderData) CancelOrder(idOrder, idFromToken int) (row int, err error) {
	dataOrder := Order{}
	idCheck := od.db.First(&dataOrder, idOrder)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataOrder.UserID != idFromToken {
		return -1, errors.New("you don't have access")
	}
	result := od.db.Model(&Order{}).Where("id = ?", idOrder).Update("status", "Canceled")

	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update order status")
	}
	return int(result.RowsAffected), nil
}

func (od *orderData) SelectData(limit, offset int) (data []domain.Order, err error) {
	dataOrder := []Order{}
	result := od.db.Preload("Product").Preload("User").Find(&dataOrder)

	if result.Error != nil {
		return []domain.Order{}, result.Error
	}

	return ParseToArr(dataOrder), nil
}
