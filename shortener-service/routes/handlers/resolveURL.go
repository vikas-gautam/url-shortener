package handlers

import (
	"fmt"
	"net/http"
	"shortener-service/storage/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ResolveURL(c *gin.Context) {

	getParamID := c.Param("url")

	fmt.Println("printing getParamID: ", getParamID)

	actual_url, err := db.GetUrlByid(getParamID)

	if err != nil {
		logrus.Errorf("error while fetching data from db: %s", err)
		c.JSON(404, err.Error())
		return
	}
	fmt.Println("printing actual url: ", actual_url)

	c.Redirect(http.StatusTemporaryRedirect, actual_url)

}
