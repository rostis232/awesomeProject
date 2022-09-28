package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the home page handler.
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")
}

// RenderTemplate
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template ", err)
		return
	}

}
