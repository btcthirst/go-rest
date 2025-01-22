package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Start() {
	server()
}

func server() {
	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}
