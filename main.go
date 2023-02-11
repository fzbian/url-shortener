package main

import (
	routes "url-shortener/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "views")
	e.GET("/generate/:short", routes.Short)
	api := e.Group("/api")
	api.GET("/", routes.Health)
	api.POST("/addlink/", routes.AddLink)
	e.Logger.Fatal(e.Start(":8080"))
}
