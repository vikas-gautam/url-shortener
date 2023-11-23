package handlers

import (
	"auth-service/models"
	"auth-service/storage/db"
	"auth-service/storage/redis"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var Repo *Service

func NewRepo(inputService *Service) {
	Repo = inputService
}

type Service struct {
	RedisStore redis.RedisStore
	DbStore    db.DBStore
}

func (s *Service) Signup(c *gin.Context) {
	validate := validator.New()

	var userSignup models.User

	if err := c.BindJSON(&userSignup); err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("after unmarshal %v", userSignup)

	hash, _ := HashPassword(userSignup.Password)

	fmt.Println("Hash:    ", hash)

	if err := validate.Struct(userSignup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "provided data is not valid",
		})
		return
	}

	_, err := s.DbStore.InsertUser(models.DBUser{
		FirstName: userSignup.FirstName,
		LastName:  userSignup.LastName,
		Email:     userSignup.Email,
		Password:  hash,
	})
	if err != nil {
		fmt.Printf("error while saving into the db: %s", err)
		c.JSON(http.StatusInternalServerError, err)
		return
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
