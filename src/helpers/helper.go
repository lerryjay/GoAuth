package helpers


import (
	"golang.org/x/crypto/bcrypt"
	"goauth/v2/src/models"
)

type FUser struct {
	*models.User
}



func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidatePasswordHash(hashpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}
