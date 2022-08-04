package delivery

import "project3/group3/domain"

type Order struct {
	ID         int `json:"id"`
	Stock      int `json:"stock"`
	TotalPrice int `json:"total_price"`
	Product    Product
	User       User
}

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"qty"`
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func FromModel(data domain.Order) Order {
	return Order{
		ID:         data.ID,
		Stock:      data.Stock,
		TotalPrice: data.TotalPrice,
		Product: Product{
			ID:          data.Product.ID,
			ProductName: data.Product.ProductName,
			Price:       data.Product.Price,
			Stock:       data.Product.Stock,
		},
		User: User{
			ID:      data.User.ID,
			Name:    data.User.Name,
			Address: data.User.Address,
		},
	}
}

func FromModelList(data []domain.Order) []Order {
	result := []Order{}
	for i := range data {
		result = append(result, FromModel(data[i]))
	}
	return result
}
