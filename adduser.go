package database

import (

	//"time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(name string, tel string) {
	DB := connect()

	//插入資料
	stmt, err := DB.Prepare("INSERT user SET name=?,tel=?")
	checkErr(err)

	res, err := stmt.Exec(name, tel)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Printf("%d rows affected", affect)
}
