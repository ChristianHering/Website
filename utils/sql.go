package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //SQL Driver for MySQL/MariaDB
	"github.com/pkg/errors"
)

var connection *sql.DB

//Connects to our mysql db
func setupSQL(channel chan error) {
	var err error

	connection, err = sql.Open("mysql", Config.SQLConfig.Username+":"+Config.SQLConfig.Password+"@tcp("+Config.SQLConfig.Host+")/"+Config.SQLConfig.Database+"?parseTime=true")
	if err != nil {
		channel <- errors.WithStack(err)
	}
	defer connection.Close()

	go RowCountUpdater(60)

	channel <- nil
	<-make(chan struct{})
}

//RowCountUpdater asynchronously udpates row counts
func RowCountUpdater(interval int) {
	var err error

	for {
		BlogRowCount, err = GetTableSize("posts") //For GetPostRange()
		if err != nil {
			var e = Error{Date: time.Now(), Error: fmt.Sprintf("%+v", errors.WithStack(err)), Host: "", URL: ""}

			e.Create()
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

//GetTableSize returns the number of uuid's in a table
func GetTableSize(table string) (length int, err error) {
	r := connection.QueryRow(`SELECT COUNT(id) FROM ` + table + `;`)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	err = r.Scan(&length)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	err = r.Err()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return length, nil
}
