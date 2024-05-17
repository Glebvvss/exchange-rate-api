package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"time"
)

func DbOpen() (*sql.DB, error) {
	link := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", link)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}

func Migrate() error {
	db, connErr := DbOpen()
	if connErr != nil {
		return connErr
	}

	driver, driverErr := mysql.WithInstance(db, &mysql.Config{})
	if driverErr != nil {
		return driverErr
	}

	m, migrateErr := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"mysql",
		driver,
	)

	if migrateErr != nil {
		return migrateErr
	}

	m.Up()
	return nil
}
