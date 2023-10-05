package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var LOGIN_SERVICE_URL = "http://localhost:8082/api/v1/"

func Login(c *gin.Context) {

	username, password, _ := c.Request.BasicAuth()

	POST_URL := LOGIN_SERVICE_URL + "login"

	r, err := http.NewRequest("POST", POST_URL, nil)
	if err != nil {
		panic(err)
	}
	r.SetBasicAuth(username, password)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//sen the response from service as it is
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var responseData any
	json.Unmarshal(b, &responseData)
	c.JSON(res.StatusCode, responseData)

}
