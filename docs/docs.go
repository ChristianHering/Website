package docs

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves a website for public
//(and private) API documentation
func Run(m *mux.Router) {
	mux := m.Host("docs.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle("/", middlewares.ThenFunc(myHandler))

	return
}

func myHandler(w http.ResponseWriter, r *http.Request) {
}
