package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ChristianHering/Website/utils"
)

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

func StatisticsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Println(r.Host, r.URL.String()) //TODO
		fmt.Println(time.Since(start))      //TODO log this to mysql
	})
}
