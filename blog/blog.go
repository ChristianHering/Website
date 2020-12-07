package blog

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves a development blog on the blog subdomain
func Run(m *mux.Router) {
	mux := m.Host("blog.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler)

	mux.Handle("/", middlewares.ThenFunc(myHandler))

	return
}

func myHandler(w http.ResponseWriter, r *http.Request) {
}
