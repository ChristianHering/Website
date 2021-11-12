package templates

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

//Templates stores the data for all of our html templates
//Along with our template function associations
var Templates *template.Template

//Run initializes all templates held in the templates
//directory and stores them in our Templates var
func Run() error {
	templs, err := filepath.Glob("./utils/templates/*.html")
	if err != nil {
		return errors.WithStack(err)
	}

	Templates, err = template.New(templs[0]).Funcs(templateFunctions).ParseFiles(templs[0]) //Calling this in the loop with 'Templates' instead of 'template' dereferences 'Templates' when it == nil
	if err != nil {
		return errors.WithStack(err)
	}

	for i := 1; i < len(templs); i++ {
		Templates, err = Templates.New(templs[i]).Funcs(templateFunctions).ParseFiles(templs[i]) //Initialize all the html templates in the templates folder
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

//These functions are exposed to our html templates
var templateFunctions = template.FuncMap{
	"htmlSafe": func(html string) template.HTML {
		return template.HTML(html)
	},
	"increment": func(n int) int {
		return n + 1
	},
	"decrement": func(n int) int {
		return n - 1
	},
	"modulous": func(n int, x int) int {
		return n % x
	},
	"formatDate": func(n time.Time) string {
		return n.Format("Jan 2 2006")
	},
}
