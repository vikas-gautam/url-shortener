package handlers

import (
	"auth-service/models"
	"auth-service/storage/db"
	"auth-service/storage/redis"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func ResetPassword(c *gin.Context) {
	// validate := validator.New()

	resetToken := c.Param("token")

	var userInput models.User

	if err := c.BindJSON(&userInput); err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("after unmarshal %v", userInput)

	newPasswd := userInput.Password

	hashPasswd, _ := HashPassword(newPasswd)

	fmt.Printf("newPasswd and resetToken are: %v, %v: ", hashPasswd, resetToken)

	// Fetch data from Redis
	username, err := redis.GetData(resetToken)
	if err != nil {
		logrus.Error(err)
		return
	}

	//logic to check if user already exists or not
	_, err = db.GetUserByEmailid(username)
	if err != nil && err == sql.ErrNoRows {
		logrus.Error(err)
		return
	}

	//logic to  update user's passwd

	err = db.UpdateUser(username, hashPasswd)
	if err != nil {
		logrus.Error(err)
		return
	}

	// Fetch data from Redis
	err = redis.DelKey(resetToken)
	if err != nil {
		logrus.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: fmt.Sprintln("user's passwd has been reset"),
	})
}

func tokenGenerator() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func GenerateResetToken(c *gin.Context) {
	// validate := validator.New()

	var passwdReset models.User

	if err := c.BindJSON(&passwdReset); err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("after unmarshal %v", passwdReset)

	userName := passwdReset.Email

	fmt.Println("userName:    ", userName)

	generatedToken := tokenGenerator()

	fmt.Println(generatedToken)

	//inserting the same in redis
	err := redis.SetData(generatedToken, userName)
	if err != nil {
		logrus.Errorf("error while setting key in redis: %s", err)
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: fmt.Sprintf("token generated: %v", generatedToken),
	})
}
