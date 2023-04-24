package database

import (
	_ "github.com/go-sql-driver/mysql"
)

func GetUserList() []User {
	DB := connect()

	var user User
	var users []User

	//查詢資料
	rows, err := DB.Query("SELECT id, name, tel FROM user")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Tel)
		checkErr(err)
		users = append(users, user)
	}

	return users
}
