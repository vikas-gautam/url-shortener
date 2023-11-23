package handlers

import (
	"auth-service/internal"
	"auth-service/models"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var gw_reset_base = "http://localhost:9090/api/v1/reset/"

func (s *Service) ResetPassword(c *gin.Context) {
	// validate := validator.New()

	resetToken := c.Param("token")

	var userInput models.User

	if err := c.BindJSON(&userInput); err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("after unmarshal %v", userInput)

	newPasswd := userInput.Password

	hashPasswd, _ := HashPassword(newPasswd)

	fmt.Printf("newPasswd and resetToken are: %v, %v: ", hashPasswd, resetToken)

	// Fetch data from Redis

	username, err := s.RedisStore.GetData(resetToken)
	if err != nil {
		logrus.Error(err)
		return
	}

	//logic to check if user already exists or not

	_, err = s.DbStore.GetUserByEmailid(username)
	if err != nil && err == sql.ErrNoRows {
		logrus.Error(err)
		return
	}

	//logic to  update user's passwd

	err = s.DbStore.UpdateUser(username, hashPasswd)
	if err != nil {
		logrus.Error(err)
		return
	}

	// Fetch data from Redis
	err = s.RedisStore.DelKey(resetToken)
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

func (s *Service) GenerateResetToken(c *gin.Context) {
	// validate := validator.New()

	var passwdReset models.User

	if err := c.BindJSON(&passwdReset); err != nil {
		logrus.Error(err)
		return
	}

	fmt.Printf("after unmarshal %v", passwdReset)

	userName := passwdReset.Email

	fmt.Println("userName:    ", userName)

	generatedToken := tokenGenerator()

	fmt.Println(generatedToken)

	// inserting the same in redis
	err := s.RedisStore.SetData(generatedToken, userName)
	if err != nil {
		logrus.Errorf("error while setting key in redis: %s", err)
		return
	}

	// defer close(app.MailChan)
	// fmt.Println("Starting mail listner")
	// internal.ListenForMail()

	msg := models.MailData{
		From:    "mail@urlhortener.com",
		To:      userName,
		Subject: "Reset password link",
		Content: gw_reset_base + generatedToken,
	}

	internal.SendMsg(msg)

	c.IndentedJSON(http.StatusOK, models.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: fmt.Sprintf("reset password link has been sent to your email address: %v", userName),
	})
}
