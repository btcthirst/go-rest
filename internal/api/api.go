package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/btcthirst/go-rest/internal/config"
	"github.com/btcthirst/go-rest/internal/database"
	"github.com/labstack/echo/v4"
)

func Start() {
	conn := config.GetDBConn()
	e := echo.New()

	database.Init(conn)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%v", port)
	e.GET("/ping", ping)
	e.Logger.Fatal(e.Start(addr))
}

func ping(c echo.Context) error {
	now := time.Now()

	return c.JSON(http.StatusOK, now)
}
