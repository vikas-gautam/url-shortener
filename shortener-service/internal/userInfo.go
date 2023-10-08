package internal

import (
	"database/sql"
	"errors"
	"shortener-service/models"
	"shortener-service/storage/db"

	"github.com/sirupsen/logrus"
)

func UserInfo(username string) (models.DBUser, error) {

	//logic to check if user already exists or not
	userData, err := db.GetUserByEmailid(username)
	if err != nil && err == sql.ErrNoRows {
		logrus.Error(err)
		return models.DBUser{}, errors.New("user name does not exist")
	}

	// match := CheckPasswordHash(password, userData.Password)

	return userData, nil

}
