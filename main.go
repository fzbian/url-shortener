package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	routes "url-shortener/routes"
)

func main() {
	e := echo.New()
	e.GET("/", routes.Health)
	e.GET("addlink/", routes.AddLink)
	e.GET("short/:short", routes.Short)
	e.Logger.Fatal(e.Start(":8080"))

}
