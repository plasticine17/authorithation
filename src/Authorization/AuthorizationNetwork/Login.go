package AuthorizationNetwork

import (
	"database/sql"
	"Authorization/User"
	"Utils"
)

func ValidateUserLogIn(database *sql.DB, login string, password string) (bool, error) {
	
	var dbPassword, err = User.GetUserPasswordHash(database, login)
	
	//TODO: Доделать. Проверить конкретный тип ошибки
	if err != nil {
		return false, err
	}
	
	var equals = Utils.CompareHashAndPassword(dbPassword, []byte(password))
	
	return equals == nil, nil
}
