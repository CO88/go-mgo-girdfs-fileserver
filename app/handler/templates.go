package handler

import (
	"fmt"
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles("upload.html"))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", "")
    if err != nil {
    	fmt.Println("error occurred")
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func LoadUploadPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "upload")
}