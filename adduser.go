package database

import (
	"database/sql"
	"fmt"

	//"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser() {
	db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Printf("%d rows affected", affect)
}
