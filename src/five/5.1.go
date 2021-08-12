package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}


func initDB() {
	db, err := sql.Open("mysql", "root:micbay911@tcp(127.0.0.1:3306)/escollect?charset=utf8mb4")
	checkErr(err)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?, department=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("Jin", "研发部门", "2012-08-29")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("Bob", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	row, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for row.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = row.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid, username, department, created)
	}

	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	fmt.Println(affect)

	_ = db.Close()
}

func main() {
	initDB()
}