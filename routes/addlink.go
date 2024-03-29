package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"url-shortener/database"
	"url-shortener/utils"
)

type Links struct {
	short string
	link  string
}

func AddLink(ctx echo.Context) error {
	link := ctx.QueryParam("link")
	short := utils.GenString(10)

	if link == "" {
		err := ctx.String(http.StatusOK, "the param url is empty")
		if err != nil {
			return err
		}
	} else if utils.IsUrl(link) == true {
		db := database.Connect()

		_, err := db.Exec("INSERT INTO links VALUES(?, ?)", short, link)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}

		result := fmt.Sprintf("%s successfully created with the short %s\n", link, short)
		fmt.Println(result)
		res := fmt.Sprintf("Link shortened: %s", short)
		return ctx.JSON(http.StatusOK, echo.Map{
			"response": res,
		})
	} else {
		res := fmt.Sprintf("The link is invalid")
		return ctx.JSON(http.StatusOK, echo.Map{
			"response": res,
		})
	}
	return nil
}
