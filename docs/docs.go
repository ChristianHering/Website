package main

import (
	"github.com/ChristianHering/utils/middleware"
	"github.com/gorilla/mux"
)

func Run(m *mux.Router) error {
	mux := m.Host("docs.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle(middlewares.Then(myHandler))

	return nil
}
