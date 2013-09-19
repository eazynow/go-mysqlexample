package main

import (
	"database/sql"
	"fmt"
	_ "mysql"
)

type TelephonyProfile struct {
	id                string
	companyName       string
	currentPackage_id string
	email             string
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

	singleRowById(db, "12345678-EROL-AAAA-AAAA-123456789012")

}

func allRowsAsVariables(db *sql.DB) {
	fmt.Println("All rows loaded into variables")
	rows, err := db.Query("select id, email from telephony_profile")
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
	rows, err := db.Query("select id, companyName, currentPackage_id, email from telephony_profile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var profile TelephonyProfile
		rows.Scan(&profile.id, &profile.companyName, &profile.currentPackage_id, &profile.email)
		fmt.Println(profile)
	}
	rows.Close()
}

func singleRowById(db *sql.DB, profileId string) {
	fmt.Println("Single row by id")
	stmt, err := db.Prepare("select id, companyName, currentPackage_id, email from telephony_profile where id=?")
	if err != nil {
		fmt.Println(err)
		return
	}

	var profile TelephonyProfile
	err = stmt.QueryRow(profileId).Scan(&profile.id, &profile.companyName, &profile.currentPackage_id, &profile.email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(profile)

}
