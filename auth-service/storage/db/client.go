package db

import (
	"auth-service/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/sirupsen/logrus"
)

var counts int64

type DbInfo struct {
	dbClient *sql.DB
}

func openDB(dsn string) (DbInfo, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return DbInfo{}, err
	}

	err = db.Ping()
	if err != nil {
		return DbInfo{}, err
	}

	return DbInfo{
		dbClient: db,
	}, nil

}

func NewdbConnection(appConfig config.Config) (DbInfo, error) {
	dsn := appConfig.DSN
	fmt.Println(dsn)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			logrus.Warnln("Postgres is not ready yet...", err)
			counts++

		} else {
			logrus.Infof("Connected to postgres")
			return connection, nil
		}

		if counts > 10 {
			logrus.Error(err)
			return connection, err
		}

		logrus.Infof("Backing off for 2 seconds ..")
		time.Sleep(2 * time.Second)
		continue

	}
}
