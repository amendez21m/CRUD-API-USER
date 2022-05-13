package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "tester:secret@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Printf("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()
