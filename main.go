package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	println("Hello, world.")
	e.GET("/api/", handler)
	e.Logger.Fatal(e.Start(":8080"))
}

func handler(ctx echo.Context) error {
	ahan := &Ahan{
		Takagi: "takagi",
	}

	fmt.Println(ahan)
	return ctx.JSON(http.StatusOK, ahan)
}

type Ahan struct {
	Takagi string `json:"takagi"`
}
