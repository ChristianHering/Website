package portfolio

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/ChristianHering/Website/utils/templates"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves a personal portfolio website
//for showcasing the projects I've created
func Run(m *mux.Router) {
	mux := m.Host("portfolio.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler)

	mux.Handle("/", middlewares.ThenFunc(indexHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "portfolioIndex.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}
