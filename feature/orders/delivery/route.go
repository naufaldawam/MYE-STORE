package delivery

import (
	"project3/group3/domain"
	_middleware "project3/group3/feature/common"

	"github.com/labstack/echo/v4"
)

func RouteCart(e *echo.Echo, do domain.OrderHandler) {
	e.PUT("/orders/confirm", do.ConfirmStatus(), _middleware.JWTMiddleware())
	e.PUT("/orders/cancel", do.CancelStatus(), _middleware.JWTMiddleware())
}
