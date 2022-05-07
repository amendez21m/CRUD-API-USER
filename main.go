package main

import (
	"net/http"
	"users/handlers"

	"github.com/gorilla/mux"
)

func main() {
	//database.Connect()
	//CREAR TABLA
	//database.CreateTable(models.UserSchema, "users")
	//CREAMOS USERS
	//models.CreateUser("Das", "1234", "Das@gmail.es")
	//models.CreateUser("Anton", "1234", "anton@gmail.es")
	//models.CreateUser("Tere", "1234", "Tere@gmail.es")
	//database.Close()

	mux := mux.NewRouter()
	//EndPoints
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("DELETE")

	http.ListenAndServe(":3000", mux)
}
