package db

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

const (
	MYSQL = "mysql"
)

type DBConfig struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
	Net      string `json:"net"`
}

func NewClient(config DBConfig) (interface{}, error) {
	addr := config.Host + ":" + string(config.Port)

	cfg := mysql.Config{
		User:   config.User,
		Passwd: config.Secret,
		Addr:   addr,
		Net:    config.Net,
		DBName: config.Name,
	}
	db, err := sqlx.Connect(MYSQL, cfg.FormatDSN())
	if err != nil {
		fmt.Println("db not connected")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging")
		return nil, err
	}
	dbConn = db
	return dbConn, nil
}

func GetDBClient() *sqlx.DB {
	return dbConn
}
