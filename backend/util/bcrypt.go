package util

import "golang.org/x/crypto/bcrypt"

//HashPassword hash password to a hash string with cost is 14
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//VerifyPassword check password is valid
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
