package controllers

import (
	"authexample/logging"
	"authexample/routes"
	"authexample/users"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	myControllerBase
}

func (u *UserController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users.GetAllUsers())
}

func (u *UserController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser users.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	value, err := newUser.AddUser()
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(value)
	} else {
		logging.GetLogger().Info(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
}

func (u *UserController) GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["value"]
	_user, err := users.GetUser(username)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(_user)
	} else {
		logging.GetLogger().Info(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
}

func (u *UserController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["value"]
	_user, err := users.DeleteUser(username)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(_user)
	} else {
		logging.GetLogger().Info(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
}

func (u *UserController) GetRouteModel() routes.RouteConfig {
	return &users.UserRouteModel{}
}
