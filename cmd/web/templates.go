package main

import (
	"html/template"
	"path/filepath"
	"time"

	"moe.akari.best/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
}

func newTemplateCache() (map[string]*template.Template, error) {

	// Initializing a new map as in memory cache

	cache := map[string]*template.Template{}

	//get all the templates

	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	//loop through the pages

	for _, page := range pages {

		name := filepath.Base(page)

		// // Parse the base template file into a template set.
		// ts, err := template.ParseFiles("./ui/html/base.html")

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		// Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		// Call ParseFiles() *on this template set* to add the page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add the template set to the map, and use name as key

		cache[name] = ts
	}

	return cache, nil

}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}
