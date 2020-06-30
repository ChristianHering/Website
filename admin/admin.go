package admin

import (
	"fmt"
	"net/http"

	"github.com/ChristianHering/Website/utils"
	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves an administrative panel for viewing
//stats and managing content on the admin subdomain
func Run(m *mux.Router) {
	mux := m.Host("admin.christianhering.com").Subrouter()

	//Non-Authenticated Handlers

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle("/", middlewares.ThenFunc(indexHandler))

	mux.Handle("/login", middlewares.ThenFunc(utils.LoginHandler))
	mux.Handle("/logout", middlewares.ThenFunc(utils.LogoutHandler))
	mux.Handle("/callback", middlewares.ThenFunc(utils.CallbackHandler))

	//Authenticated Handlers

	middlewaresWithAuth := alice.New(middleware.ErrorHandler, middleware.AuthenticationHandler, middleware.StatisticsHandler)

	mux.Handle("/dashboard", middlewaresWithAuth.ThenFunc(dashboardHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}
