package database

import (
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	email string
	username string
	passwordhash string
	fullname string
	createDate string
	role int
 }

var userList = []user{
	{
		email:        "abc@gmail.com",
		username:     "abc12",
		passwordhash: "hashedme1",
		fullname:     "abc def",
		createDate:   "1631600786",
		role:         1,
	},
	{
		email:        "chekme@example.com",
		username:     "checkme34",
		passwordhash: "hashedme2",
		fullname:     "check me",
		createDate:   "1631600837",
		role:         0,
	},
}

func GetUserObject(email string) (user, bool) {
	//needs to be replaces using Database
	for _, user := range userList {
		if user.email == email {
			return user, true
		}
	}
	return user{}, false
}

// checks if the password hash is valid
func (u *user) ValidatePasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.passwordhash), []byte(password))
	return err == nil
}

// this simply adds the user to the list
func AddUserObject(email string, username string, password string, fullname string, role int) bool {
	// declare the new user object
	hashpassword, err := HashPassword(password)
	if err != nil {
		return false
	}

	newUser := user{
		email:        email,
		passwordhash: hashpassword,
		username:     username,
		fullname:     fullname,
		role:         role,
	}
	// check if a user already exists
	for _, ele := range userList {
		if ele.email == email || ele.username == username {
			return false
		}
	}
	userList = append(userList, newUser)
	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}