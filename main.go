package main

import (
	"github.com/labstack/echo/v4"
	"github.com/webtaken/go-templ/handler"
)

type DB struct{}

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	// userHandler := handler.UserHandler{}
	e.GET("/", handler.Home)
	e.GET("/signup", handler.SignUp)
	e.Logger.Fatal(e.Start(":3001"))
}
