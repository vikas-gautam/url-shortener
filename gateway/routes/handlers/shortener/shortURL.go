package shortener

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var SHORTENER_SERVICE_URL = "http://localhost:8081/api/v1/"

func ShortURL(c *gin.Context) {

	username, password, _ := c.Request.BasicAuth()
	body, _ := io.ReadAll(c.Request.Body)

	POST_URL := SHORTENER_SERVICE_URL + "shorturl"

	r, err := http.NewRequest("POST", POST_URL, bytes.NewBuffer(body))
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

	//send the response from service as it is
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var responseData any
	json.Unmarshal(b, &responseData)

	c.JSON(res.StatusCode, responseData)
}
