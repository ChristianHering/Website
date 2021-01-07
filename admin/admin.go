package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ChristianHering/Website/utils"
	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/ChristianHering/Website/utils/templates"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves an administrative panel for viewing
//stats and managing content on the admin subdomain
func Run(m *mux.Router) {
	mux := m.Host("admin.christianhering.com").Subrouter()

	//Non-Authenticated Handlers

	middlewares := alice.New(middleware.ErrorHandler)

	mux.Handle("/", middlewares.ThenFunc(indexHandler))

	mux.Handle("/login", middlewares.ThenFunc(utils.LoginHandler))
	mux.Handle("/logout", middlewares.ThenFunc(utils.LogoutHandler))
	mux.Handle("/callback", middlewares.ThenFunc(utils.CallbackHandler))

	//Authenticated Handlers

	middlewaresWithAuth := alice.New(middleware.ErrorHandler, middleware.AuthenticationHandler)

	mux.Handle("/dashboard", middlewaresWithAuth.ThenFunc(dashboardHandler))

	mux.Handle("/error", middlewaresWithAuth.ThenFunc(errorHandler))
	mux.Handle("/errorDelete", middlewaresWithAuth.ThenFunc(errorDeletionHandler))
	mux.Handle("/errorDeleteType", middlewaresWithAuth.ThenFunc(errorTypeDeletionHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "adminDashboard.html", nil)
	if err != nil {
		panic(err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	var e utils.Errors

	err := e.Read("3")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	json.NewEncoder(w).Encode(e)
}

func errorDeletionHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var e utils.Error

	err = json.Unmarshal(b, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	err = e.Delete()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func errorTypeDeletionHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var e utils.Error

	err = json.Unmarshal(b, &e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	err = e.DeleteErrorType()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
