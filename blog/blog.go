package blog

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func Run(m *mux.Router) error {
	mux := m.Host("blog.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle("/", middlewares.ThenFunc(myHandler))
	return nil
}

func myHandler(w http.ResponseWriter, r *http.Request) {
}
