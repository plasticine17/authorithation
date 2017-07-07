package AuthorizationNetwork

import (
	"database/sql"
)

//Регистрирует нового пользователя
func RegisterNewUser(database *sql.DB, info RegisterInfo) error {
	return registerNewUser(database, info)
}


