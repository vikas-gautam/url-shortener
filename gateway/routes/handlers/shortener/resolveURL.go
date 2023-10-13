package shortener

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// var RESOLVE_SERVICE_URL = "http://localhost:8081/"

func ResolveURL(c *gin.Context) {

	getParamID := c.Param("url")

	GET_URL := os.Getenv("RESOLVE_SERVICE_URL") + getParamID

	req, err := http.NewRequest("GET", GET_URL, nil)
	if err != nil {
		panic(err)

	}

	client := &http.Client{
		CheckRedirect: func(redirect *http.Request, via []*http.Request) error {
			fmt.Println("Redirected to:", redirect.URL)
			return http.ErrUseLastResponse
		}}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 307 {
		//send the response from service as it is
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var responseData any
		json.Unmarshal(b, &responseData)
		c.JSON(resp.StatusCode, responseData)
		return
	}
	redirect_location, _ := resp.Location()

	c.Redirect(http.StatusTemporaryRedirect, redirect_location.String())

}
