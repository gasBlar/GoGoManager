package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlUrl() string {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "gogomanager"

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

	return db, nil
}

func Init() (*sql.DB, error) {
	dbSource := GetMysqlUrl()

	return NewDB("mysql", dbSource)
}
