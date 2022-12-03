package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(ctx echo.Context) error {
	indexMessage := "I'm actually ok!"
	return ctx.String(http.StatusOK, indexMessage)
}
