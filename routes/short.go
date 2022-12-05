package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"url-shortener/database"
)

func Short(ctx echo.Context) error {
	short := ctx.Param("short")
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

	for _, bk := range bks {
		for {
			if bk.short == short {
				err := ctx.Redirect(http.StatusPermanentRedirect, bk.link)
				if err != nil {
					return err
				}
				break
			} else {
				err := ctx.String(http.StatusOK, "The short is not valid")
				if err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}
