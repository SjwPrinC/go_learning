package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//mysql connection schame: user:pwd@(host:port)/db
	db, _ := sql.Open("mysql", "root:sjw202020220@(114.215.41.183:13306)/testdb")

	err := db.Ping()

	if err != nil {
		fmt.Println("connect failed", err.Error())
		return
	}
	defer db.Close()

	QueryTest(db)

	// InsertTest(db)

	// UpdateTest(db)

	// DeleteTest(db)
}

func QueryTest(db *sql.DB) {
	var u user
	r, _ := db.Query("select * from user")
	for r.Next() {

		r.Scan(&u.id, &u.name, &u.age)

		fmt.Println(u)
	}
}

func InsertTest(db *sql.DB) {
	sqlStr := "insert into user(name,age) values(?,?)"
	result, err := db.Exec(sqlStr, "hongdan", 30)

	if err != nil {
		fmt.Println("insert failed", err.Error())
	} else {
		fmt.Println(result)
	}

}

func UpdateTest(db *sql.DB) {
	result, err := db.Exec("update user set age = ? where id = ?", 31, 1)

	if err != nil {
		fmt.Println("update failed", err.Error())
	} else {
		fmt.Println(result)
	}
}

func DeleteTest(db *sql.DB) {

}

type user struct {
	id   int
	name string
	age  int
}
