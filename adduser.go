package database

import (

	//"time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Print("adduser")
	// connect()
}

func Snd() {

}

// func AddUser(name string, tel string) {
// 	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
// 	checkErr(err)

// 	//插入資料
// 	stmt, err := db.Prepare("INSERT userinfo SET name=?,tel=?")
// 	checkErr(err)

// 	res, err := stmt.Exec("astaxie", "研發部門")
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Printf("%d rows affected", affect)
// }
