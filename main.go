package main

import (
	"github.com/labstack/echo/v4"

	database "url-shortener/database"
	routes "url-shortener/routes"
)

func main() {
	e := echo.New()
	e.GET("/", routes.Index)
	e.GET("addlink/", routes.AddLink)
	database.Connect()
	e.Logger.Fatal(e.Start(":8080"))
}
