// PackAge main is an example of how to connect to a mysql database
// using golang, and a few different examples of reading data
package main

import (
	"database/sql"
	"fmt"
	_ "mysql"
)

type User struct {
	Id    string
	Name  string
	Age   int
	Email string
}

// String representation of the user
func (u User) String() (s string) {
	s := fmt.Sprintf("User id=%s, Name=%s, Age=%s, Email=%s", u.Id, u.Name, u.Age, u.Email)
	return
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

// allRowsAsVariables just reads data from the users table and stored the fields
// in local variables. Useful for ad hoc queries where you may not have a struct
// defined for the response
func allRowsAsVariables(db *sql.DB) {
	fmt.Println("All rows loaded into variables")
	rows, err := db.Query("select id, Email from users")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var Email string
		rows.Scan(&id, &Email)
		fmt.Println(id, Email)
	}
	rows.Close()
}

// allRowsAsStruct returns reads the data from the users table and loads each
// row into a struct. This is the more likely way to get back data as you have
// defined the structure in advance
func allRowsAsStruct(db *sql.DB) {
	fmt.Println("All rows loaded into a struct")
	rows, err := db.Query("select id, name, Age, Email from users")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		fmt.Println(user)
	}
	rows.Close()
}

// singleRowById shows an example of looking up a row. It uses parameters to avoid
// sql injection, and loads the response into a struct
func singleRowById(db *sql.DB, id string) {
	fmt.Println("Single row by id")
	stmt, err := db.Prepare("select id, name, Age, Email from users where id=?")
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Name, &user.Email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

}
