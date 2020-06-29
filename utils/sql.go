package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
)

var session *gocql.Session

//Connects to our cassandra cluster
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

var pageAdverageLatency map[string]float64 = make(map[string]float64)
var pageLatencyCount map[string]int = make(map[string]int)

var siteAdverageLatency map[string]float64 = make(map[string]float64)
var siteLatencyCount map[string]int = make(map[string]int)

//Calculates the average latency for a request's page
//and site, then stores the results in a hashmap
func LogPageLatency(latency time.Duration, host string, url string) {
	count := siteLatencyCount[host] //Latency statistics for a request's subdomain
	siteAdverageLatency[host] = (siteAdverageLatency[host]*float64(count) + latency.Seconds()) / float64(count+1)
	siteLatencyCount[host] = count + 1

	count = pageLatencyCount[host+url] //Latency statistics for a request's page
	pageAdverageLatency[host+url] = (pageAdverageLatency[host+url]*float64(count) + latency.Seconds()) / float64(count+1)
	pageLatencyCount[host+url] = count + 1
}
