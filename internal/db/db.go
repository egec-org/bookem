package db

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

type DBConfig struct {
	Host   string
	Name   string
	User   string
	Passwd string
}

func NewClient(config DBConfig) (interface{}, error) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "127.0.0.1:3306",
		DBName:               config.Name,
		AllowNativePasswords: true,
	}
	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
		fmt.Println("db not connected")
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
		fmt.Println("Error pinging")
	}
	dbConn = db
	return dbConn, nil
}

func GetDBClient() *sqlx.DB {
	return dbConn
}
