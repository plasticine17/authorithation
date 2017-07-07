package Authorization

import (
	"Authorization/AuthorizationNetwork"
	"database/sql"
)

type Executor struct {
	Database *sql.DB
}


func (e Executor) Register(info AuthorizationNetwork.RegisterInfo) error {
	return AuthorizationNetwork.RegisterNewUser(e.Database, info)
}

func (e Executor) Login(login string, password string) (bool, error) {
	return AuthorizationNetwork.ValidateUserLogIn(e.Database, login, password)
}
