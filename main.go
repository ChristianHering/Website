package main

import (
	"fmt"
	"net/http"

	"github.com/ChristianHering/admin"
	"github.com/ChristianHering/blog"
	"github.com/ChristianHering/consulting"
	"github.com/ChristianHering/docs"
	"github.com/ChristianHering/portfolio"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func main() {
	err := utils.runConfigSetup() //Utility initialization routine
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.WithStack(err)))
	}

	mux := mux.NewRouter()

	admin.Run(mux)      //admin subdomain
	blog.Run(mux)       //blog subdomain
	consulting.Run(mux) //top level domain
	docs.Run(mux)       //docs subdomain
	portfolio.Run(mux)  //portfolio subdomain

	err = http.ListenAndServe(":80", mux)
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.WithStack(err)))
	}
}
