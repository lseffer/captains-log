package controllers

import (
	"captains-log/db"
	"captains-log/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// LoginController handles logging in
func LoginController(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	connection, err := db.GetDatabaseConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	returnUser, err := user.Read(connection)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !user.Password.Valid || !returnUser.PasswordHash.Valid {
		return
	}
	passwordHash := returnUser.PasswordHash.String
	password := user.Password.String
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err == nil {
		sessionTokenID := uuid.New().String()
		sessionTokenValidTo := time.Now().Add(120 * time.Second)
		session := models.Session{ID: sessionTokenID, UserID: returnUser.ID, ValidTo: sessionTokenValidTo.Unix()}
		err1 := session.DeleteByUser(connection, returnUser)
		err2 := session.Create(connection)
		if err1 != nil || err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "X-Session-Token",
			Value:   sessionTokenID,
			Expires: sessionTokenValidTo,
		})
		json.NewEncoder(w).Encode(returnUser)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// LogoutController handles logging out
func LogoutController(w http.ResponseWriter, req *http.Request) {
}
