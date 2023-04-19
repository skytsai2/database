package database

import (
	"fmt"
)

func init() {
	fmt.Print("checkerr")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
