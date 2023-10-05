package handlers

import (
	"auth-service/internal"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "error parsing basic auth",
		})
		return
	}

	authenticated, _, err := internal.Auth(username, password)

	if !authenticated {
		c.IndentedJSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, models.Response{
		Status:  http.StatusText(http.StatusOK),
		Message: "User has successfully authenticated",
	})

}
