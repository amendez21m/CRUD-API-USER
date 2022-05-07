package models

import (
	"fmt"
	"users/database"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	fmt.Println("createuser")
	return user
}

func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := database.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, err := database.Query(sql)
	if err != nil {
		fmt.Println("rompo")
	}
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

func GetUser(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users where id=?"
	rows, _ := database.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func (user *User) Update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	database.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.Update()
	}
}

func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	database.Exec(sql, user.Id)
}
