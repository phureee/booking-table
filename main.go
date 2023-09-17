package main

import (
	"booking-table/routers"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routers.InitRouter(e)
	e.Logger.Fatal(e.Start(":8080"))
}
