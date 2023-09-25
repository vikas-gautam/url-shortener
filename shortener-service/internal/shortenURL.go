package internal

import (
	"log"
	"shortener-service/models"
	"shortener-service/storage/db"

	"github.com/google/uuid"
)

// actual logic to shorten the url
func ShortenURL(actual_url string, userData models.DBUser) (string, error) {

	//logic to shorten the actual url in the request payload
	id := uuid.New().String()[:5]

	newMapping := models.DBURL{UserID: userData.ID, ActualURL: actual_url, ShortURL: id}

	//save the mapping into the database
	GeneratedId, err := db.InsertUrl(newMapping)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return GeneratedId, nil

}
