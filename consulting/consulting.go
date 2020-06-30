package consulting

import (
	"fmt"
	"net/http"

	"github.com/ChristianHering/Website/utils/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//Run Serves our main website
func Run(m *mux.Router) {
	mux := m.Host("christianhering.com").Subrouter()

	middlewares := alice.New(middleware.ErrorHandler, middleware.StatisticsHandler)

	mux.Handle("/", middlewares.ThenFunc(myHandler))

	return
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!!")
	fmt.Println("hi")
}
