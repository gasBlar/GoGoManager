package db

import (
	"database/sql"
	"fmt"

	"github.com/gasBlar/GoGoManager/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetMysqlUrl() string {
	dbHost := config.GetEnv("MYSQL_HOST")
	dbPort := config.GetEnv("MYSQL_PORT")
	dbUser := config.GetEnv("MYSQL_USER")
	dbPass := config.GetEnv("MYSQL_PASSWORD")
	dbName := config.GetEnv("MYSQL_DATABASE")

	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

func NewDB(dbDriver string, dbSource string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	m, err := migrate.New(
		"file://db/migration",
		"mysql://"+dbSource,
	)

	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No new migration to apply")
			return db, nil
		} else {
			fmt.Println("Error migrating the database: ", err)
			return nil, err
		}
	} else {
		fmt.Println("Migration successful")
		return db, nil
	}
}

func InitDb() (*sql.DB, error) {
	dbSource := GetMysqlUrl()

	return NewDB("mysql", dbSource)
}
