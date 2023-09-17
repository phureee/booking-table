package routers

import (
	"booking-table/service"

	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "service ok")
	})

	e.POST("/table", service.InitTable)
	e.GET("/table", service.GetTable)

	e.POST("/booking", service.BookingTable, CheckInitTable)
	e.GET("/booking", service.GetBook)

	e.POST("/booking/cancel", service.CancelBooking, CheckInitTable)
}
