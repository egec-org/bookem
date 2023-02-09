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
	Protocol string `json:"protocol"`
}

func NewClient(config DBConfig) (interface{}, error) {
	addr := config.Host + ":" + config.Port
	fmt.Println(addr)

	cfg := mysql.Config{
		User:                 config.User,
		Passwd:               config.Secret,
		Addr:                 addr,
		Net:                  config.Protocol,
		DBName:               config.Name,
		AllowNativePasswords: true,
	}
	fmt.Println(cfg)
	db, err := sqlx.Connect(MYSQL, cfg.FormatDSN())
	if err != nil {
		fmt.Println("db not connected")
		fmt.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	dbConn = db
	return dbConn, nil
}

func GetDBClient() *sqlx.DB {
	return dbConn
}
