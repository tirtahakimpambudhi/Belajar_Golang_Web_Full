package controller

import (
	"html/template"
	"net/http"
)
var templates , _ = template.ParseGlob("./templates/*.gohtml")
func TemplateFileGlob (w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w,"index.gohtml","Hello World")
} 

func TemplateFile(w http.ResponseWriter, r *http.Request) {
	t , err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	t.ExecuteTemplate(w,"index.gohtml","Home Page")
}