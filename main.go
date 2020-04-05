package main

import (
	"captains-log/controllers"
	"captains-log/db"
	"flag"
	"log"
	"net/http"
)

func main() {
	autoApplyRevisions := flag.Bool("a", false, "automatically attempt to apply revisions on startup")

	flag.Parse()

	if *autoApplyRevisions {
		db.ApplyRevisions()
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.IndexController)

	log.Println("Server listening on port 5555")
	log.Fatal(http.ListenAndServe(":5555", nil))
}
