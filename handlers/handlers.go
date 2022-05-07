package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"users/database"
	"users/models"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	database.Connect()
	users := models.ListUsers()
	database.Close()

	output, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	database.Connect()
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.GetUser(userId)
	database.Close()
	output, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)

	} else {
		database.Connect()
		user.Save()
		database.Close()
	}
	output, _ := json.Marshal(user)
	fmt.Fprintf(rw, string(output))

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)

	} else {
		database.Connect()
		user.Save()
		database.Close()
	}
	output, _ := json.Marshal(user)
	fmt.Fprintf(rw, string(output))

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	database.Connect()
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.GetUser(userId)
	user.Delete()
	database.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintf(rw, string(output))
}
