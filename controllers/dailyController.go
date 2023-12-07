package controllers

import (
	"database/sql"
	"fmt"
	"time"
)

func Daily(db *sql.DB) (string, error) {
	var picked_on string
	today := time.Now().Format("2006-01-02")

	err := db.QueryRow("SELECT picked_on FROM daily where picked_on = $1", today).Scan(&picked_on)
	if err != nil {
		//Assume there no daily was picked yet for the day
		fmt.Println(err)
	}

	return picked_on, err
}

func pick(db *sql.DB) {

}
