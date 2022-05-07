package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "tester:secret@tcp(localhost:3306)/test"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connecting...")
	db = connection
}

func Close() {
	db.Close()
	fmt.Println("Disconnecting...")

}

func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Success")
		}
	} else {
		fmt.Println("This table already exist.")
	}
}

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
