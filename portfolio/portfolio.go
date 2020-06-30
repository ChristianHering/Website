package portfolio

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves a personal portfolio website
//for showcasing the projects I've created
func Run(m *mux.Router) {
	mux := m.Host("portfolio.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle("/", middlewares.ThenFunc(myHandler))

	return
}

func myHandler(w http.ResponseWriter, r *http.Request) {
}
