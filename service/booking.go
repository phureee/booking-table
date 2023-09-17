package service

import (
	"booking-table/entities"
	"booking-table/pkg/response"
	"fmt"
	"time"

	"github.com/labstack/echo"
)

type bookRequest struct {
	CustomerAmount int `json:"customer_amount"`
}

type bookRespone struct {
	BookID       int   `json:"booking_id"`
	TablesBooked int   `json:"tables_booked"`
	TablesID     []int `json:"tables_id"`
}

func BookingTable(c echo.Context) error {
	var bookReq bookRequest

	if err := c.Bind(&bookReq); err != nil {
		fmt.Println("Bind err", err.Error())
		return err
	}

	if bookReq.CustomerAmount <= 0 {
		return response.Fail(c, 400, "customer amount should more than 0")
	}

	amountCust := bookReq.CustomerAmount

	var bookedTablesID []int
	var bookSuccessed bool

	tables := entities.GetTable()
	for tableID, table := range tables {
		if table.IsAvailable {
			bookedTablesID = append(bookedTablesID, tableID)
			amountCust = amountCust - table.SeaterAmount
			if amountCust <= 0 {
				bookSuccessed = true
				break
			}
		}
	}

	if !bookSuccessed {
		return response.Fail(c, 400, "not enough tables for all customers")
	}

	for _, tableID := range bookedTablesID {
		entities.UpdateTableAvailable(tableID, false)
	}

	bookID := entities.SetBook(entities.Booking{
		CustomerAmount: bookReq.CustomerAmount,
		TableID:        bookedTablesID,
		IsAvailable:    true,
		CreatedAt:      time.Now(),
	})

	return response.Success(c, 200, bookRespone{
		BookID:       bookID,
		TablesBooked: len(bookedTablesID),
		TablesID:     bookedTablesID,
	})
}

func GetBook(c echo.Context) error {
	return response.Success(c, 200, entities.ShowBook())
}
