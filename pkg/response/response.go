package response

import "github.com/labstack/echo"

func Success(e echo.Context, status int, message interface{}) error {
	res := make(map[string]interface{})

	res["status"] = status

	switch message.(type) {
	case string:
		res["message"] = message
	default:
		res["result"] = message
	}

	return e.JSON(200, res)
}

func Fail(e echo.Context, status int, message interface{}) error {
	return e.JSON(400, map[string]interface{}{
		"status":  status,
		"message": message,
	})
}
