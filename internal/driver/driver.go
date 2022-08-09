package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConne = 5 // might be changed in production
const maxIdleDbConne = 5 // might be changed in production
const maxDbLifeTime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConne)
	d.SetMaxIdleConns(maxIdleDbConne)
	d.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}
	fmt.Println("*** Pinged database successfully! ***")

	return nil
}
