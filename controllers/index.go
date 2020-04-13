package controllers

import (
	"captains-log/models"
	"html/template"
	"net/http"
)

// IndexController renders index.html
func IndexController(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseGlob("views/*"))
	dummyData := []models.Post{
		{ID: 1, Title: "", Text: "something", TimeCreated: 123456789},
		{ID: 2, Title: "", Text: "something1", TimeCreated: 123456789},
		{ID: 3, Title: "", Text: "something23", TimeCreated: 123456789},
	}
	t.ExecuteTemplate(w, "index.html", dummyData)
}
