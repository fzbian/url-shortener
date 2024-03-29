package routes

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"url-shortener/database"
)

func Short(ctx echo.Context) error {
	short := ctx.Param("short")
	if short == "favicon.ico" {
		return nil
	}
	db := database.Connect()

	rows, err := db.Query("SELECT * FROM links")
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	bks := make([]*Links, 0)
	for rows.Next() {
		bk := new(Links)
		err := rows.Scan(&bk.short, &bk.link)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Panic(err.Error())
		return nil
	}

	var link string
	found := false
	for _, bk := range bks {
		if bk.short == short {
			link = bk.link
			found = true
			break
		}
	}
	if found {
		err := ctx.Redirect(http.StatusPermanentRedirect, link)
		if err != nil {
			return err
		}
	} else {
		err := ctx.File("views/404.html")
		if err != nil {
			return err
		}
	}
	return nil
}
