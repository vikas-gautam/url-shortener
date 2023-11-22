package handlers

import (
	"auth-service/models"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(c *gin.Context) {

	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "error parsing basic auth",
		})
		return
	}

	authenticated, _, err := s.EnsureAuth(username, password)

	if !authenticated {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: "User has successfully logged in",
	})

}

//###############  internal package ################

func (s *Service) EnsureAuth(username, password string) (bool, models.DBUser, error) {

	//logic to check if user already exists or not
	userData, err := s.Store.GetUserByEmailid(username)
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
