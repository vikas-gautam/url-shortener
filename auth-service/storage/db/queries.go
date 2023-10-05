package db

import (
	"auth-service/models"
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func Connection(conn *sql.DB) {
	db = conn
}

// Insert a new user into the database, and returns the ID of the newly inserted row
func InsertUser(mapping models.DBUser) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	query := `insert into users (first_name, last_name, email, password, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	err := db.QueryRowContext(ctx, query,
		mapping.FirstName,
		mapping.LastName,
		mapping.Email,
		mapping.Password,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func GetUserByEmailid(email string) (models.DBUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, first_name, last_name, email, password, is_active, created_at, updated_at from users where email = $1`

	var data models.DBUser

	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&data.ID,
		&data.FirstName,
		&data.LastName,
		&data.Email,
		&data.Password,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return data, err
	}

	return data, nil
}
