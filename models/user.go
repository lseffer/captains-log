package models

import (
	"database/sql"

	"gopkg.in/guregu/null.v3"
)

// User model for a user
type User struct {
	ID           uint32      `json:"id" db:"id"`
	PasswordHash null.String `json:"password_hash" db:"password_hash"`
	Password     null.String `json:"password"`
}

// Read user from database
func (u *User) Read(connection *sql.DB) (User, error) {
	var user User
	err := connection.QueryRow("SELECT * FROM users WHERE id=?", u.ID).Scan(&user.ID, &user.PasswordHash, &user.Password)
	if err != nil {
		return user, err
	}
	return user, err
}

// Create user from database
func (u *User) Create(connection *sql.DB) error {
	_, err := connection.Exec("INSERT INTO users VALUES(?, ?, ?)", u.ID, u.PasswordHash, u.Password)
	return err
}

// Update user from database
func (u *User) Update(connection *sql.DB) error {
	_, err := connection.Exec("UPDATE users SET password=:password password_hash=:password_hash WHERE id=:id", sql.Named("id", u.ID), sql.Named("password", u.Password), sql.Named("password_hash", u.PasswordHash))
	return err
}

// Delete user from database
func (u *User) Delete(connection *sql.DB) error {
	_, err := connection.Exec("DELETE FROM users WHERE id=?", u.ID)
	return err
}
