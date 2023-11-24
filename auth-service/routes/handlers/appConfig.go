package handlers

import (
	"auth-service/storage/db"
	"auth-service/storage/redis"
)

var Repo *Service

func NewRepo(inputService *Service) {
	Repo = inputService
}

type Service struct {
	RedisStore redis.RedisStore
	DbStore    db.DBStore
}
