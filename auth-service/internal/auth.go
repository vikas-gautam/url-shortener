package internal

import (
	"auth-service/models"
	"auth-service/storage/db"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Auth(username, password string) (bool, models.DBUser, error) {

	//logic to check if user already exists or not
	userData, err := db.GetUserByEmailid(username)
	if err != nil && err == sql.ErrNoRows {
		logrus.Error(err)
		return false, models.DBUser{}, errors.New("user name does not exist")
	}

	match := CheckPasswordHash(password, userData.Password)

	//email not exists
	if !match {
		return false, models.DBUser{}, errors.New("invalid credentials")
	} else {
		return true, userData, nil
	}

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
