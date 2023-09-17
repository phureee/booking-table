package service

import (
	"booking-table/entities"
	"booking-table/pkg/response"
	"time"

	"github.com/labstack/echo"
)

type tableRequest struct {
	TableAmount          int `json:"table_amount"`
	SeaterAmountPerTable int `json:"seater_per_table"`
}

func InitTable(c echo.Context) error {
	var tableReq tableRequest
	if err := c.Bind(&tableReq); err != nil {
		return err
	}

	if entities.AmountTable() > 0 {
		return response.Fail(c, 400, "init table only once")
	}

	if tableReq.TableAmount <= 0 {
		return response.Fail(c, 400, "table amount should more than 0")
	}

	for i := 0; i < tableReq.TableAmount; i++ {
		seatAmount := 4
		if tableReq.SeaterAmountPerTable > 0 {
			seatAmount = tableReq.SeaterAmountPerTable
		}
		tableID := i + 1
		entities.SetTable(tableID, entities.Table{
			SeaterAmount: seatAmount,
			IsAvailable:  true,
			CreatedAt:    time.Now(),
		})
	}

	return response.Success(c, 201, "initialized table success")
}

func GetTable(c echo.Context) error {
	return response.Success(c, 200, entities.ShowTable())
}
