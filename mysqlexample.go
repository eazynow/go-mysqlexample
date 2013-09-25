// Package main is an example of how to connect to a mysql database
// using golang, and a few different examples of reading data
package main

import (
	"database/sql"
	"fmt"
	_ "mysql"
)

type User struct {
	id    string
	name  string
	age   int
	email string
}

func main() {
	username := "user1"
	password := "mypass"
	host := "localhost:3306"
	schema := "test"
	charset := "utf8"

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", username, password, host, schema, charset)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// a bit like finally clause
	defer db.Close()

	allRowsAsVariables(db)
	allRowsAsStruct(db)

	singleRowById(db, "USER0001")

}

func allRowsAsVariables(db *sql.DB) {
	fmt.Println("All rows loaded into variables")
	rows, err := db.Query("select id, email from users")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var email string
		rows.Scan(&id, &email)
		fmt.Println(id, email)
	}
	rows.Close()
}

func allRowsAsStruct(db *sql.DB) {
	fmt.Println("All rows loaded into a struct")
	rows, err := db.Query("select id, name, age, email from users")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.id, &user.name, &user.age, &user.email)
		fmt.Println(user)
	}
	rows.Close()
}

func singleRowById(db *sql.DB, id string) {
	fmt.Println("Single row by id")
	stmt, err := db.Prepare("select id, name, age, email from users where id=?")
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User
	err = stmt.QueryRow(id).Scan(&user.id, &user.name, &user.name, &user.email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

}
