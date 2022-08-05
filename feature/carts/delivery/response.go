package delivery

import "project3/group3/domain"

type Cart struct {
	ID         int `json:"id"`
	Stock      int `json:"stock"`
	TotalPrice int `json:"total_price"`
	UserID     int `json:"user_id"`
	Product    Product
}

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

func FromModel(data domain.Cart) Cart {
	return Cart{
		ID:         data.ID,
		Stock:      data.Stock,
		TotalPrice: data.TotalPrice,
		UserID:     data.UserID,
		Product: Product{
			ID:          data.Product.ID,
			ProductName: data.Product.ProductName,
			Price:       data.Product.Price,
			Stock:       data.Product.Stock,
		},
	}
}

func FromModelList(data []domain.Cart) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromModel(data[key]))
	}
	return result
}
