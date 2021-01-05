package consulting

import (
	"net/http"
	"strings"

	"github.com/ChristianHering/Website/utils"
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
	mux.Handle("/contact", middlewares.ThenFunc(contactHandler))

	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.Templates.ExecuteTemplate(w, "consultingIndex.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}

//Only requires an email
func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	var subject string
	var body string

	if r.Form["message"] != nil {
		subject = "Web Contact - " + strings.Join(r.Form["name"], "")
		body = strings.Join(r.Form["message"], "")
	} else {
		subject = "Web Follow Up Request"
		body = strings.Join(r.Form["email"], "")
	}

	var e utils.Email = utils.Email{
		From:    strings.Join(r.Form["email"], ""),
		To:      "Contact@ChristianHering.com",
		Subject: subject,
		Body:    body,
	}

	err = e.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	err = utils.SendMail(&e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	http.Redirect(w, r, "http://ChristianHering.com/", http.StatusSeeOther)
}
