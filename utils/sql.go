package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var connection *sql.DB

//Connects to our mysql db
func setupSQL(channel chan error) {
	var err error

	connection, err = sql.Open("postgres", "user="+Config.SQLConfig.Username+" password="+Config.SQLConfig.Password+" host="+Config.SQLConfig.Host+" port="+Config.SQLConfig.Port+" sslmode=disable")
	if err != nil {
		channel <- errors.WithStack(err)
	}
	defer connection.Close()
	connection.Exec(`set search_path='` + Config.SQLConfig.Schema + `'`)

	channel <- nil
	<-make(chan struct{})
}
