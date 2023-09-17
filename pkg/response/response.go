package response

import "github.com/labstack/echo"

func Success(e echo.Context, status int, message interface{}) error {
	return e.JSON(200, map[string]interface{}{
		"status":  status,
		"message": message,
	})
}

func Fail(e echo.Context, status int, message interface{}) error {
	return e.JSON(400, map[string]interface{}{
		"status":  status,
		"message": message,
	})
}
