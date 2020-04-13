//+build !testsuite

package main

import (
	"captains-log/controllers"
	"captains-log/db"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	autoApplyRevisions := flag.Bool("a", false, "automatically attempt to apply revisions on startup")

	flag.Parse()

	if *autoApplyRevisions {
		db.Startup()
	}

	r := mux.NewRouter().StrictSlash(true)

	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(controllers.AuthMiddleWare)

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	protectedRoutes.HandleFunc("/", controllers.IndexController).Methods(http.MethodGet)
	r.HandleFunc("/auth/login", controllers.LoginController).Methods(http.MethodPost)
	r.HandleFunc("/auth/logout", controllers.LogoutController).Methods(http.MethodPost)

	protectedRoutes.HandleFunc("/posts/create", controllers.PostController).Methods(http.MethodGet)
	protectedRoutes.HandleFunc("/posts/{id}", controllers.PostController).Methods(http.MethodGet)
	protectedRoutes.HandleFunc("/posts/{id}", controllers.PostController).Methods(http.MethodDelete)
	protectedRoutes.HandleFunc("/posts/{id}", controllers.PostController).Methods(http.MethodPut)
	protectedRoutes.HandleFunc("/posts/{id}", controllers.PostController).Methods(http.MethodPost)

	log.Println("Server listening on port 5555")
	log.Fatal(http.ListenAndServe(":5555", r))
}
