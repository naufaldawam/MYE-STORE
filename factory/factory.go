package factory

import (
	ud "project3/group3/feature/users/data"
	userDelivery "project3/group3/feature/users/delivery"
	us "project3/group3/feature/users/usecase"

	pd "project3/group3/feature/products/data"
	productDelivery "project3/group3/feature/products/delivery"
	ps "project3/group3/feature/products/usecase"

	cd "project3/group3/feature/carts/data"
	cartDelivery "project3/group3/feature/carts/delivery"
	cs "project3/group3/feature/carts/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := ud.New(db)
	useCase := us.New(userData, validator)
	userHandler := userDelivery.New(useCase)
	userDelivery.RouteUser(e, userHandler)

	productData := pd.New(db)
	productCase := ps.New(productData)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteProduct(e, productHandler)

	cartData := cd.New(db)
	cartCase := cs.New(cartData)
	cartHandler := cartDelivery.New(cartCase)
	cartDelivery.RouteCart(e, cartHandler)
}
