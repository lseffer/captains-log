package models

import "database/sql"

// Session ...
// main model for a user session
type Session struct {
	ID      string `json:"token_id"`
	UserID  uint32 `json:"user_id"`
	ValidTo int64  `json:"valid_to"`
}

// Read session from database
func (u *Session) Read(connection *sql.DB) (Session, error) {
	var session Session
	err := connection.QueryRow("SELECT * FROM sessions WHERE id=?", u.ID).Scan(&session.ID, &session.UserID, &session.ValidTo)
	if err != nil {
		return session, err
	}
	return session, err
}

// Create session from database
func (u *Session) Create(connection *sql.DB) error {
	_, err := connection.Exec("INSERT INTO sessions VALUES(?, ?, ?)", u.ID, u.UserID, u.ValidTo)
	return err
}

// Update session from database
func (u *Session) Update(connection *sql.DB) error {
	_, err := connection.Exec("UPDATE sessions SET user_id=:user_id valid_to=:valid_to WHERE id=:id", sql.Named("id", u.ID), sql.Named("user_id", u.UserID), sql.Named("valid_to", u.ValidTo))
	return err
}

// Delete session from database
func (u *Session) Delete(connection *sql.DB) error {
	_, err := connection.Exec("DELETE FROM sessions WHERE id=?", u.ID)
	return err
}

// DeleteByUser session from database
func (u *Session) DeleteByUser(connection *sql.DB, user User) error {
	_, err := connection.Exec("DELETE FROM sessions WHERE user_id=?", user.ID)
	return err
}
