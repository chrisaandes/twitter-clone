package utils

import "golang.org/x/crypto/bcrypt"

/*PasswordGenerator ...*/
func PasswordGenerator(password string) (string, error) {
	hard := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hard)
	return string(bytes), err
}
