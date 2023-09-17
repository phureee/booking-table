package service_test

import (
	"booking-table/service"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func BookingTable(e *echo.Echo, customerAmount int) {

	SetTable(e, 1)
	requestJSON := fmt.Sprintf(`{"customer_amount": %d}`, customerAmount)

	req := httptest.NewRequest(http.MethodPost, "/booking", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service.BookingTable(c)
}

func TestCancelBooking(t *testing.T) {
	// Setup
	e := echo.New()

	BookingTable(e, 4)

	requestJSON := `{ "booking_id": 1}`

	req := httptest.NewRequest(http.MethodPost, "/booking/cancel", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectResponse := `{"result":{"amount_of_cancel":1,"tables_id":[1]},"status":200}`
	// Assertions
	if assert.NoError(t, service.CancelBooking(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expectResponse, rec.Body.String())
	}
}

func TestCancelBookingErrorWrongBookID(t *testing.T) {
	// Setup
	e := echo.New()

	BookingTable(e, 4)

	requestJSON := `{ "booking_id": 2}`

	req := httptest.NewRequest(http.MethodPost, "/booking/cancel", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectResponse := `{"message":"booking id not found", "status":404}`
	// Assertions
	if assert.NoError(t, service.CancelBooking(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, expectResponse, rec.Body.String())
	}
}
