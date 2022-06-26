package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/michaeldeven/tokodeven/api/database"
	"github.com/michaeldeven/tokodeven/api/models"
	"github.com/michaeldeven/tokodeven/api/repository"
	"github.com/michaeldeven/tokodeven/api/repository/crud"
	"github.com/michaeldeven/tokodeven/api/responses"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		users, err := usersRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err = usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err := usersRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
