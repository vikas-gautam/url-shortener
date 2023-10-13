package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		logrus.Info("The request   is started \n")

		username, password, _ := c.Request.BasicAuth()

		GET_URL := os.Getenv("AUTH_SERVICE_URL") + "auth"

		r, err := http.NewRequest("GET", GET_URL, nil)
		if err != nil {
			panic(err)
		}

		r.SetBasicAuth(username, password)
		r.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {

			//sen the response from service as it is
			b, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			var responseData any
			json.Unmarshal(b, &responseData)
			c.JSON(res.StatusCode, responseData)
			c.Abort()
			return
		}
		c.Next()
		logrus.Info("The request is served \n")
	}
}
