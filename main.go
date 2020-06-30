package main

import (
	"fmt"
	"net/http"

	"github.com/ChristianHering/Website/admin"
	"github.com/ChristianHering/Website/blog"
	"github.com/ChristianHering/Website/cdn"
	"github.com/ChristianHering/Website/consulting"
	"github.com/ChristianHering/Website/docs"
	"github.com/ChristianHering/Website/portfolio"
	"github.com/ChristianHering/Website/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func main() {
	err := utils.RunUtilSetup() //Utility initialization routine
	if err != nil {
		panic(err)
	}

	mux := mux.NewRouter()

	cdn.Run(mux) //cdn subdomain

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
