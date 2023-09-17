package routers

import (
	"booking-table/entities"
	"booking-table/pkg/response"

	"github.com/labstack/echo"
)

func CheckInitTable(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		if entities.AmountTable() == 0 {
			return response.Fail(c, 400, "init table before")
		}
		return next(c)
	}
}
