package database

import (

	//"time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func EditUser(userid int, tel string) {
	DB := connect()

	//更新資料
	stmt, err := DB.Prepare("update user set tel=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(tel, userid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
