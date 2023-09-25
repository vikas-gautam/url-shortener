package models

type Response struct {
	Status  string
	Message string
}


type ShortenerResponse struct {
	ActualURL  string
	ShortenedURL string
}