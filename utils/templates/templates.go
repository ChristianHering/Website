package templates

import (
	"html/template"
)

var Templates *template.Template

//Initializes all templates held in the templates
//directory and stores them in our Templates var
func Run() {
	Templates = template.Must(template.ParseGlob("./utils/templates/*.html")) //Initialize all the html templates in the templates folder
}
