package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}


func initDBSqlite() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkError(err)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	stmt, err := db.Prepare("INSERT INTO userinfo (username, department, created) VALUES (?,?,?)")
	checkError(err)

	res, err := stmt.Exec("Jin", "研发部门", "2012-08-29")
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkError(err)

	res, err = stmt.Exec("Bob", id)
	checkError(err)

	affect, err := res.RowsAffected()
	checkError(err)
	fmt.Println(affect)

	row, err := db.Query("SELECT * FROM userinfo")
	checkError(err)

	for row.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = row.Scan(&uid, &username, &department, &created)
		checkError(err)

		fmt.Println(uid, username, department, created)
	}

	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkError(err)

	res, err = stmt.Exec(id)
	checkError(err)

	affect, err = res.RowsAffected()
	fmt.Println(affect)

	_ = db.Close()
}

func main() {
	initDBSqlite()
}
