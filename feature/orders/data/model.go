package data

import (
	"project3/group3/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Stock     int
	Status    string
	UserID    int
	ProductID int
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Role     string
	Address  string
	Product  []Product
}

type Product struct {
	gorm.Model
	Stock        int
	Status       string
	ProductName  string
	ProductImage string
	Price        int
	UserID       int
	User         User
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		ID:        int(o.ID),
		Stock:     o.Stock,
		Status:    o.Status,
		UserID:    o.UserID,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
		Product: domain.ProductOrder{
			ID:          int(o.Product.ID),
			ProductName: o.Product.ProductName,
			Stock:       o.Product.Stock,
			Price:       o.Product.Price,
		},
	}
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order
	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func FromDomain(data domain.Order) Order {
	var res Order
	res.Stock = data.Stock
	res.Status = data.Status
	res.UserID = data.UserID
	res.ProductID = data.Product.ID
	return res
}
