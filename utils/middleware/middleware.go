package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func ErrorHandler(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				err = fmt.Sprintf("%+v", errors.WithStack(err.(error)))
				log.Println(err) //TODO log this to mysql
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

func StatisticsHandler(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Println(time.Since(start)) //TODO log this to mysql
	}

	return http.HandlerFunc(f)
}
