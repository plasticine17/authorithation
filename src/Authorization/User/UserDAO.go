package User

import (
	"database/sql"
)

const (
	modelTable = "main.users"

	modelColums = "identifier, login, email, password_hash, name_alias, registration_id"

	insertSQL = "INSERT INTO " + modelTable + " (" + modelColums + ") VALUES ($1, $2, $3, $4, $5, $6)"

	updateSQL = "UPDATE " + modelTable + " SET regID = $1 WHERE identifier = $2"

	selectUserByLoginSQL = "SELECT " + modelColums + " FROM " + modelTable + " WHERE login = $1"

	existsLoginValidationSQL = "SELECT EXISTS(SELECT 1 FROM " + modelTable + " WHERE login = $1)"
)



func (u Model) Insert(database *sql.DB) error {
	var _, err = database.Exec(insertSQL, u.ID.String(), u.Login, u.Email, string(u.PasswordHash), u.NameAlias, u.RegistrationID)

	return err
}

func (u Model) UpdateRegistrationID(database *sql.DB, newRegistrationID string) error {
	var _, err = database.Exec(updateSQL, newRegistrationID, u.ID.String())

	return err
}

func SelectUserByLogin(database *sql.DB, login string) (Model, error) {
	var m Model
	var passwordHash []byte

	var err = database.QueryRow(selectUserByLoginSQL, login).Scan(&m.ID, &m.Login, &m.Email, &passwordHash, &m.NameAlias, &m.RegistrationID)

	if err != nil {
		return m, err
	}

	m.PasswordHash = passwordHash

	return m, nil
}

func GetUserPasswordHash(database *sql.DB, login string) ([]byte, error) {
	var model, err = SelectUserByLogin(database, login)

	return model.PasswordHash, err
}

func IsLoginExists(database *sql.DB, login string) (bool, error) {
	var exists bool
	var err = database.QueryRow(existsLoginValidationSQL, login).Scan(&exists)

	return exists, err
}
