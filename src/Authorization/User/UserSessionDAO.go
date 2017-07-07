package User

import (
	"database/sql"
	"github.com/satori/go.uuid"
)

const (
	insertSessionSQL = "INSERT INTO " + sessionTable + " (" + sessionColumns + ") VALUES ($1, $2, $3, $4)"

	sessionColumns = "identifier, userid, logintime, lastseentime"

	sessionTable = "auth.session"

	validateSessionSQL = "SELECT EXISTS(SELECT 1 FROM " + sessionTable + " WHERE session = $1 AND userid = $2)"
)

func (s Session) Insert(database sql.DB) error {
	var _, err = database.Exec(insertSessionSQL, s.SessionID, s.UserID, s.LoginTime, s.LastSeenTime)
	return err
}

func ValidateSession(database *sql.DB, session string, userID uuid.UUID) (bool, error) {

	var isValid bool

	var err = database.QueryRow(validateSessionSQL, session, userID.String()).Scan(&isValid)

	return isValid, err
}

