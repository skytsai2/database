package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	UserName     string = "root"
	Password     string = "root"
	Addr         string = "mysql"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func connect() *sql.DB {
	//組合sql連線字串
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	//連接MySQL
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return DB
	}
	DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	DB.SetMaxOpenConns(MaxOpenConns)
	DB.SetMaxIdleConns(MaxIdleConns)

	return DB
}
