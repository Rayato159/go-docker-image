package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type message struct {
	Msg string `json:"msg"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		h := &message{
			Msg: "hello, Docker 🐳",
		}
		return c.JSON(http.StatusOK, h)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
