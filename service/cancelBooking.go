package service

import (
	"booking-table/entities"
	"booking-table/pkg/response"

	"github.com/labstack/echo"
)

type cancelRequest struct {
	BookingID int `json:"booking_id"`
}

type cancelRespone struct {
	CancelAmount int   `json:"amount_of_cancel"`
	TablesID     []int `json:"tables_id"`
}

func CancelBooking(c echo.Context) error {
	var cancelReq cancelRequest
	if err := c.Bind(&cancelReq); err != nil {
		return err
	}

	if cancelReq.BookingID <= 0 {
		return response.Fail(c, 400, "booking id is require")
	}

	booking := entities.GetBook()

	book, ok := booking[cancelReq.BookingID]
	if !ok {
		return response.Fail(c, 404, "booking id not found")
	}

	if !book.IsAvailable {
		return response.Fail(c, 400, "this booking id have canceled")
	}

	entities.UpdateBookingAvailable(cancelReq.BookingID, false)

	for _, id := range book.TableID {
		entities.UpdateTableAvailable(id, true)
	}

	return response.Success(c, 200, cancelRespone{
		CancelAmount: len(book.TableID),
		TablesID:     book.TableID,
	})
}
