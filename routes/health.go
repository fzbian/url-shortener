package routes

import (
	"fmt"
	"log"
	"net/http"

	"url-shortener/database"

	"github.com/labstack/echo/v4"
)

func Health(ctx echo.Context) error {

	_, err := database.PingDB()
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Println("\nHealth\nServer is Ok!\nDatabase is Ok!")
	}
	indexMessage := "Ok!"
	return ctx.String(http.StatusOK, indexMessage)
}
