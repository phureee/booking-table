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

func SetTable(e *echo.Echo, tableAmount int) {
	requestJSON := fmt.Sprintf(`{"table_amount": %d}`, tableAmount)

	req := httptest.NewRequest(http.MethodPost, "/table", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	service.InitTable(c)
}

func TestBooking(t *testing.T) {
	// Setup
	e := echo.New()

	SetTable(e, 1)

	requestJSON := `{"customer_amount": 3}`

	req := httptest.NewRequest(http.MethodPost, "/booking", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectResponse := `{"result":{"booking_id":1,"tables_booked":1,"tables_id":[1]},"status":200}`
	// Assertions
	if assert.NoError(t, service.BookingTable(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expectResponse, rec.Body.String())
	}
}

func TestBookingErrorNotEnoughTables(t *testing.T) {
	// Setup
	e := echo.New()

	SetTable(e, 1)

	requestJSON := `{"customer_amount": 5}`

	req := httptest.NewRequest(http.MethodPost, "/booking", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectResponse := `{"message":"not enough tables for all customers", "status":400}`
	// Assertions
	if assert.NoError(t, service.BookingTable(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, expectResponse, rec.Body.String())
	}
}
