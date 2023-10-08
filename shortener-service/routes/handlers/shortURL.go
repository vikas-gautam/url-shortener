package handlers

import (
	"net/http"
	"shortener-service/internal"
	"shortener-service/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var GatewayURL = "http://localhost:9090"

func ShortURL(c *gin.Context) {

	var url models.URL

	if err := c.BindJSON(&url); err != nil {
		logrus.Error(err)
		return
	}
	username := c.GetHeader("username")

	userData, err := internal.UserInfo(username)
	if err != nil {
		logrus.Error(err.Error())
	}

	//very important
	actual_url := ""
	if !strings.HasPrefix(url.ActualURL, "http://") && !strings.HasPrefix(url.ActualURL, "https://") {
		actual_url = "http://" + url.ActualURL
	}

	//replcae it with middleware
	GeneratedId, _ := internal.ShortenURL(actual_url, userData)

	shortURL := GatewayURL + "/" + GeneratedId

	c.JSON(http.StatusOK, models.ShortenerResponse{
		ActualURL:    url.ActualURL,
		ShortenedURL: shortURL,
	})

}
