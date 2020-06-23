package utils

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
)

var session *gocql.Session

func setupSQL(channel chan error) {
	var err error

	cluster := gocql.NewCluster(Config.SqlConfig.Nodes)
	cluster.Keyspace = "website"
	cluster.Consistency = gocql.Quorum
	session, err = cluster.CreateSession()
	if err != nil {
		channel <- errors.WithStack(err)
	}
	defer session.Close()
	/*
		err = session.Query(`INSERT INTO errors (id, error, host, url) VALUES (?, ?, ?, ?)`, //CREATE
			gocql.TimeUUID(), err, host, url).Exec()
		if err != nil {
			panic(fmt.Sprintf("%+v", errors.WithStack(err)))
		}

		var id string
		var e string

		iter := session.Query(`SELECT id, error FROM errors WHERE host = ?`, "param1").Iter() //READ
		for iter.Scan(&id, &e) {
			fmt.Println("Tweet:", id, e)

			err = session.Query(`UPDATE errors SET error = ?, host = ? WHERE id = ?`, //UPDATE
				"error! I fucked up...", "just param1", id).Exec()
			if err != nil {
				panic(fmt.Sprintf("%+v", errors.WithStack(err)))
			}

		}
		if err := iter.Close(); err != nil {
			panic(fmt.Sprintf("%+v", errors.WithStack(err)))
		}

		iter = session.Query(`SELECT id FROM errors WHERE host = ?`, "just param1").Iter()
		for iter.Scan(&id) {
			err = session.Query(`DELETE FROM errors WHERE id = ? IF EXISTS`, id).Exec() //DELETE
			if err != nil {
				panic(fmt.Sprintf("%+v", errors.WithStack(err)))
			}
		}
		if err := iter.Close(); err != nil {
			panic(fmt.Sprintf("%+v", errors.WithStack(err)))
		}
	*/
	channel <- nil
	<-make(chan struct{})
}

//Logs an error to the errors table for later analysis
func LogError(e string, host string, url string) {
	err := session.Query(`INSERT INTO errors (id, error, host, url) VALUES (?, ?, ?, ?)`,
		gocql.TimeUUID(), e, host, url).Exec()
	if err != nil {
		log.Println("Failed to log the following error to errors DB", err)
		log.Println("Error encountered: ", fmt.Sprintf("%+v", errors.WithStack(err)))
	}
}
