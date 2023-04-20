package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DelUser(userid string) {
	DB := connect()

	//刪除資料
	stmt, err := DB.Prepare("delete from user where id=?")
	checkErr(err)

	res, err := stmt.Exec(userid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	DB.Close()
}
