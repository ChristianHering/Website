package utils

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //SQL Driver for MySQL/MariaDB
	"github.com/pkg/errors"
)

//Errors is a slice of 'Error'
type Errors []Error

//Error represents a row in the 'errors' DB table
type Error struct {
	ID    int
	Date  time.Time
	Error string
	Host  string
	URL   string
}

//Create (Err) Logs an error to the errors table
func (e *Error) Create() {
	_, err := connection.Exec(`INSERT INTO errors (date, error, host, url) VALUES (?, ?, ?, ?);`, time.Now().Format("2006-01-02 15:04:05.000000"), e.Error, e.Host, e.URL)
	if err != nil {
		log.Println("Failed to log the following error to errors DB", err)
		log.Println("Error encountered: ", fmt.Sprintf("%+v", errors.WithStack(err)))
	}
}

//Read (Err) Reads the last 'limit' errors from the error table
func (e *Errors) Read(limit string) error {
	r, err := connection.Query(`SELECT * FROM errors ORDER BY id DESC LIMIT ?;`, limit)
	if err != nil {
		return errors.WithStack(err)
	}
	defer r.Close()

	for r.Next() {
		var date string
		var er Error

		err = r.Scan(&er.ID, &date, &er.Error, &er.Host, &er.URL)
		if err != nil {
			return errors.WithStack(err)
		}

		er.Date, err = time.Parse("2006-01-02 15:04:05.000000", date)
		if err != nil {
			return errors.WithStack(err)
		}

		*e = append(*e, er)
	}

	err = r.Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//Delete (Err) Deletes an error in the error table
func (e *Error) Delete() error {
	_, err := connection.Exec(`DELETE FROM errors WHERE id = ?;`, e.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//DeleteErrorType (Err) Deletes all instances of an error in the error table
func (e *Error) DeleteErrorType() error {
	_, err := connection.Exec(`DELETE FROM errors WHERE error = ?;`, e.Error)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
