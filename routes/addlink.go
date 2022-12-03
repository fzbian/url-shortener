package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"url-shortener/utils"
)

func AddLink(ctx echo.Context) error {
	link := ctx.QueryParam("link")
	if link == "" {
		ctx.String(http.StatusOK, "the param url is empty")
	} else if utils.IsUrl(link) == true {
		ctx.String(http.StatusOK, "the url is valid")
	} else {
		ctx.String(http.StatusOK, "the url is not valid")
	}
	return nil
}
