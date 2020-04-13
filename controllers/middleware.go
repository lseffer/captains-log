package controllers

import (
	"captains-log/db"
	"captains-log/models"
	"log"
	"net/http"
	"time"
)

// AuthMiddleWare checks that session exists and is still valid
func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		connection, err := db.GetDatabaseConnection()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var reqSession models.Session
		token := r.Header.Get("X-Session-Token")
		reqSession.ID = token
		dbSession, err := reqSession.Read(connection)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid or missing token"))
			return
		}
		if dbSession.ValidTo >= time.Now().Unix() {
			log.Printf("Authenticated user %d\n", dbSession.UserID)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			log.Printf("Token expired %s\n", dbSession.ID)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token expired"))
		}
	})
}
