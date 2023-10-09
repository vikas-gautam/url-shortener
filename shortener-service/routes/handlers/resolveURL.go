package handlers

import (
	"fmt"
	"net/http"
	"shortener-service/storage/db"
	"shortener-service/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ResolveURL(c *gin.Context) {

	getParamID := c.Param("url")

	fmt.Println("printing getParamID: ", getParamID)

	// Fetch data from Redis
	actual_url, err := redis.GetData(getParamID)

	if actual_url == "" || err != nil {
		// Data not found in Redis, fetch from the database
		actual_url, err := db.GetUrlByid(getParamID)
		if err != nil {
			logrus.Errorf("error while fetching data from db: %s", err)
			c.JSON(404, err.Error())
			return
		}
		//inserting the same in redis
		err = redis.SetData(getParamID, actual_url)
		if err != nil {
			logrus.Errorf("error while setting key in redis: %s", err)
			return
		}
		logrus.Info("Cache MISS")

	} else {
		logrus.Info("Cache HIT")
	}

	fmt.Println("printing actual url: ", actual_url)

	c.Redirect(http.StatusTemporaryRedirect, actual_url)

}
