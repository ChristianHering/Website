package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ChristianHering/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
)

func Run(m *mux.Router) error {
	mux := m.Host("consulting.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle(middlewares.Then(myHandler))

	return nil
}

func myHandler(next http.Handler) http.Handler {
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
