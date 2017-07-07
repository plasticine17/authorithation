package AuthorizationNetwork

import (
	"database/sql"
	
	"Authorization/User"
	"github.com/satori/go.uuid"
	"errors"
	"Utils"
)


//Регистрирует нового пользователя
func registerNewUser(database *sql.DB, info RegisterInfo) error {
	
	var login, email, password, alias = extractInfoFromRegisterInfo(info)
	
	var alreadyExists, err0 = validateLoginExistance(database, login)
	
	if err0 != nil {
		return err0
	} else if alreadyExists {
		return errors.New("User already registred")
	}
	
	var newHashedPassword, err = Utils.CryptPassword([]byte(password))
	
	if err != nil {
		return err
	}
	
	var registrationID = uuid.NewV4()
	var newUser = User.Model{uuid.NewV4(), login, email, newHashedPassword, alias, registrationID}
	
	err = newUser.Insert(database)
	
	if err == nil {
		go sendVerification(info, registrationID)
	}
	
	return err
}

//Извлекает информацию уз структуры
//Login, Email, Password, Alias
func extractInfoFromRegisterInfo(info RegisterInfo) (string, string, string, string) {
	
	var login, email, password, alias string
	
	login = info.Login
	email = info.Email
	password = info.Password
	alias = info.NameAlias
	
	if login == "" {
		login = email
	}
	
	if alias == "" {
		alias = login
	}
	
	return login, email, password, alias
}

func validateLoginExistance(database *sql.DB, login string) (bool, error) {
	return User.IsLoginExists(database, login)
}


func sendVerification(info RegisterInfo, registrationID uuid.UUID) {
	
	switch info.RegisterType {
	
	case RegisterTypeEmail:
		sendVerificationEmail(info.Email, registrationID)
	
	case RegisterTypeTelegram:
		sendVerificationTelegramCode(info.Login)
		
	}
	
}


func sendVerificationEmail(toEmail string, registrationID uuid.UUID) {
	
	var link = "http://" + Utils.Domain + "/accounts/v1/verify/" + registrationID.String()
	var signUpHTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+ "\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"
	
	Utils.SendEmail([]string{toEmail}, "Verification", signUpHTML)
}


func sendVerificationTelegramCode(login string) {

}