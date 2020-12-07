package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //SQL Driver for MySQL/MariaDB
	"github.com/pkg/errors"
)

var connection *sql.DB

//Connects to our mysql db
func setupSQL(channel chan error) {
	var err error

	connection, err = sql.Open("mysql", Config.SQLConf.Username+":"+Config.SQLConf.Password+"@tcp("+Config.SQLConf.Host+")/"+Config.SQLConf.Database+"?parseTime=true")
	if err != nil {
		channel <- errors.WithStack(err)
	}
	defer connection.Close()

	channel <- nil
	<-make(chan struct{})
}

//LogError Logs an error to the errors table for later analysis
func LogError(e string, host string, url string) {
	_, err := connection.Query(`INSERT INTO errors (date, error, host, url) VALUES (?, ?, ?, ?);`, time.Now(), e, host, url)
	if err != nil {
		log.Println("Failed to log the following error to errors DB", err)
		log.Println("Error encountered: ", fmt.Sprintf("%+v", errors.WithStack(err)))
	}
}
