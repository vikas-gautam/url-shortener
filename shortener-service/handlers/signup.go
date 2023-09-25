package handlers

import (
	"fmt"
	"net/http"
	"shortener-service/models"

	"shortener-service/storage/db"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	validate := validator.New()

	var userSignup models.User

	if err := c.BindJSON(&userSignup); err != nil {
		log.WithFields(log.Fields{
			"userInput": userSignup,
		}).Error(err)
		return
	}

	hash, _ := HashPassword(userSignup.Password)

	fmt.Println("Hash:    ", hash)

	if err := validate.Struct(userSignup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "provided data is not valid",
		})
		return
	}

	//save the mapping into the database
	_, err := db.InsertUser(models.DBUser{
		FirstName: userSignup.FirstName,
		LastName:  userSignup.LastName,
		Email:     userSignup.Email,
		Password:  hash,
	})
	if err != nil {
		fmt.Printf("error while saving into the db: %s", err)
	}

	c.IndentedJSON(http.StatusOK, models.Response{
		Status:  http.StatusText(http.StatusAccepted),
		Message: "User signed up successfully",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
