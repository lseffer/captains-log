package main

import (
	"captains-log/controllers"
	"captains-log/db"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	db.InitializeDatabase()
	http.HandleFunc("/", controllers.IndexController)
	log.Println("Server listening on port 5555")
	log.Fatal(http.ListenAndServe(":5555", nil))
}
