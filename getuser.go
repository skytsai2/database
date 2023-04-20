package database

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Tel  string
}

func GetUser(userid string) User {
	DB := connect()

	var user User
	row := DB.QueryRow("select id, name, tel from user where id=?", userid)
	err := row.Scan(&user.ID, &user.Name, &user.Tel)
	checkErr(err)

	return user
}
