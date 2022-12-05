package database

import (
	"fmt"
)

func PingDB() (result string, error error) {
	db := Connect()

	err := error
	if err = db.Ping(); err != nil {
		var fatalMessage = fmt.Sprintf("the database did not respond, %s", err)
		fmt.Println(fatalMessage)
	} else {
		fmt.Println("the database responded well")
	}
	return result, err
}
