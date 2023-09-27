package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:staffing/out
//go:embed all:personal/public
var embeddedFS embed.FS

func main() {
	personalFS, err := fs.Sub(embeddedFS, "personal/public")
	if err != nil {
		log.Fatal(err)
	}

	staffingFS, err := fs.Sub(embeddedFS, "staffing/out")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("christianhering.com/", http.FileServer(http.FS(personalFS)))

	http.Handle("staffing.christianhering.com/", http.FileServer(http.FS(staffingFS)))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
