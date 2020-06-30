package cdn

//Despite the package name, this does
//not attempt to create or simulate a
//content delivery network. It simply
//serves static files for cdn caching

import (
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves static resources on subdomain 'cdn'
//and downloads with Cache Control headers enabled
func Run(m *mux.Router) {
	mux := m.Host("cdn.christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.CacheControlHandler, middleware.StatisticsHandler)

	mux.PathPrefix("/css/").Handler(middlewares.Then(http.StripPrefix("/css/", http.FileServer(http.Dir("./cdn/css")))))
	mux.PathPrefix("/js/").Handler(middlewares.Then(http.StripPrefix("/js/", http.FileServer(http.Dir("./cdn/js")))))

	return
}
