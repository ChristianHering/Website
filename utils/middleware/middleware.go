package middleware

import (
	"fmt"
	"net/http"

	"github.com/ChristianHering/Website/utils"
	"github.com/pkg/errors"
)

//Captures any errors that are encountered while handling a request.
//
//Sends any errors recieved to our LogError function,
//so it can be stored in our DB for future review
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				utils.LogError(fmt.Sprintf("%+v", err), r.Host, r.URL.String())
			}
		}()

		next.ServeHTTP(w, r)
	})
}

//Protects an endpoint that requires authentication.
//Redirects unauthenticated users to  our login page
func AuthenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := utils.SessionStore.Get(r, "auth-session")
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			panic(errors.WithStack(err))
		}

		if _, ok := session.Values["profile"]; !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

//Sets cache control header for all requests
func CacheControlHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, must-revalidate, proxy-revalidate, max-age="+utils.Config.MaxCacheAge))
		next.ServeHTTP(w, r)
	})
}
