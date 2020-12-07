package consulting

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/ChristianHering/Website/utils/templates"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves our main website
func Run(m *mux.Router) {
	mux := m.Host("christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler)

	mux.Handle("/", middlewares.ThenFunc(indexHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		panic(err)
	}
}
