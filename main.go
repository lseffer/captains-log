package main

import (
	"net/http"

	"captains-log/controllers"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println(controllers.Dummy(1))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
