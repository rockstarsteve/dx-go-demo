package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type user struct {
	userId   string
	userName string
	password string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1)/dx-server"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("打开数据库有问题！")
	}
	rowdata := db.QueryRow("select user_id,username, password  from sys_user where user_id = ?", "1")
	var u user
	err2 := rowdata.Scan(&u.userId, &u.userName, &u.password)
	if err2 != nil {
		fmt.Println("查询出问题了!", err2)
	}
	fmt.Println(u)
	time.Sleep(time.Second * 5)
}
