package handlers

import (
	"net/http"
	"shortener-service/internal"
	"shortener-service/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var FrontendDomain = "http://localhost:8081"

func ShortURL(c *gin.Context) {

	var url models.URL

	if err := c.BindJSON(&url); err != nil {
		logrus.Error(err)
		return
	}

	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "error parsing basic auth",
		})
		return
	}

	authenticated, userData, err := internal.Auth(username, password)

	if !authenticated {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}

	//very important
	actual_url := ""
	if !strings.HasPrefix(url.ActualURL, "http://") && !strings.HasPrefix(url.ActualURL, "https://") {
		actual_url = "http://" + url.ActualURL
	}

	GeneratedId, _ := internal.ShortenURL(actual_url, userData)

	shortURL := FrontendDomain + "/" + GeneratedId

	c.IndentedJSON(http.StatusAccepted, models.ShortenerResponse{
		ActualURL:    url.ActualURL,
		ShortenedURL: shortURL,
	})

}
