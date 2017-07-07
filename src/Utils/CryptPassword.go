package Utils

import "golang.org/x/crypto/bcrypt"

func clearMemory(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0;
	}
}

func CryptPassword(password []byte) ([]byte, error) {
	defer clearMemory(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}


func CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}