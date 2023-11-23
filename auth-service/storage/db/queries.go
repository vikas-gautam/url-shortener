package db

import (
	"auth-service/models"
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

// if not using method
// var db *sql.DB
// func Connection(conn *sql.DB) {
// 	db = conn
// }

//using DI ****************************************************************

type DBStore interface {
	InsertUser(models.DBUser) (int, error)
	GetUserByEmailid(string) (models.DBUser, error)
	UpdateUser(string, string) error
}

func NewStore(db *sql.DB) DBStore {
	return &store{db}
}

// The actual store would contain some state. In this case it's the sql.db instance, that holds the connection to our database
type store struct {
	db *sql.DB
}

//************************************************************************************************

// Insert a new user into the database, and returns the ID of the newly inserted row
func (d *store) InsertUser(mapping models.DBUser) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	query := `insert into users (first_name, last_name, email, password, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6) returning id`

	err := d.db.QueryRowContext(ctx, query,
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

func (d *store) GetUserByEmailid(email string) (models.DBUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id, first_name, last_name, email, password, is_active, created_at, updated_at from users where email = $1`

	var data models.DBUser

	row := d.db.QueryRowContext(ctx, query, email)

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

func (d *store) UpdateUser(email, newpasswd string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `update users set password = $2, updated_at = $3 where email = $1`

	_, err := d.db.ExecContext(ctx, query,
		email,
		newpasswd,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
