package handlers

import (
	"encoding/json"
	"gorm/database"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	database.Database.Find(&users)
	sendData(rw, users, http.StatusOK)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	user, err := getUserById(r)
	if err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)

	}
}

func getUserById(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	err := database.Database.First(&user, userId)

	if err.Error != nil {
		return user, err
	} else {
		return user, nil
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		database.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	var userId int64

	oldUser, err := getUserById(r)

	if err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = oldUser.Id

		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&user)
		if err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			database.Database.Save(&user)
			sendData(rw, user, http.StatusCreated)
		}
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	user, err := getUserById(r)
	if err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		database.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}

}
