package service_test

import (
	"booking-table/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestInitTable(t *testing.T) {
	// Setup
	e := echo.New()

	requestJSON := `{"table_amount": 3}`

	req := httptest.NewRequest(http.MethodPost, "/table", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectResponse := `{"message":"initialized table success","status":201}`
	// Assertions
	if assert.NoError(t, service.InitTable(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, expectResponse, rec.Body.String())
	}
}

func TestInitTableError(t *testing.T) {
	// Setup
	e := echo.New()

	requestJSON := `{"table_amount": 0}`

	req := httptest.NewRequest(http.MethodPost, "/table", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectResponse := `{"message":"table amount should more than 0","status":400}`
	// Assertions
	if assert.NoError(t, service.InitTable(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.JSONEq(t, expectResponse, rec.Body.String())
	}
}

func TestInitTableDupilcate(t *testing.T) {
	// Setup
	e := echo.New()

	requestJSON := `{"table_amount": 3}`

	req := httptest.NewRequest(http.MethodPost, "/table", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, service.InitTable(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	req2 := httptest.NewRequest(http.MethodPost, "/table", strings.NewReader(requestJSON))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	expectResponseDup := `{"message":"init table only once","status":400}`
	if assert.NoError(t, service.InitTable(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
		assert.JSONEq(t, expectResponseDup, rec2.Body.String())
	}
}
