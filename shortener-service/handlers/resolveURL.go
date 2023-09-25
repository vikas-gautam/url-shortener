package handlers

import (
	"net/http"
	"shortener-service/storage/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ResolveURL(c *gin.Context) {

	getParamID := c.Param("url")

	actual_url, err := db.GetUrlByid(getParamID)

	if err != nil {
		logrus.Errorf("error %s", err)
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, actual_url)

}
