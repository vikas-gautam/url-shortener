package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var SIGNUP_SERVICE_URL = "http://localhost:8082/api/v1/"

func Signup(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	POST_URL := SIGNUP_SERVICE_URL + "signup"

	r, err := http.NewRequest("POST", POST_URL, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

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
