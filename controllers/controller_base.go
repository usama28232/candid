package controllers

import (
	"authexample/routes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type controllerBase interface {
	GetAllHandler(http.ResponseWriter, *http.Request)
	CreateHandler(http.ResponseWriter, *http.Request)
	GetHandler(http.ResponseWriter, *http.Request)
	DeleteHandler(http.ResponseWriter, *http.Request)
	GetRouteModel() routes.RouteConfig
}

type myControllerBase struct {
}

func (cb *myControllerBase) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w)
}

func (cb *myControllerBase) CreateHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w)
}

func (cb *myControllerBase) GetHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w)
}

func (cb *myControllerBase) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w)
}

func (cb *myControllerBase) GetRouteModel() (routes.RouteConfig, error) {
	return nil, errors.New("route model not configured correctly")
}

func register(controller controllerBase, r *mux.Router) *mux.Router {
	rm := controller.GetRouteModel()
	availableRoutes := rm.Init()
	baseRoute, baseRErr := rm.GetBaseRoute()

	if baseRErr == nil && availableRoutes.AllowListAPI {
		r.HandleFunc(baseRoute, controller.GetAllHandler).Methods(http.MethodGet)
	}

	if availableRoutes.AllowAddAPI {
		addRoute, addRErr := rm.GetAddRoute(rm)
		if addRErr == nil {
			r.HandleFunc(addRoute, controller.CreateHandler).Methods(http.MethodPost)
		}
	}

	if availableRoutes.AllowDetailAPI {
		getInfoRoute, getInfoRErr := rm.GetInfoRoute(rm)
		if getInfoRErr == nil {
			r.HandleFunc(getInfoRoute, controller.GetHandler).Methods(http.MethodGet)
		}
	}

	if availableRoutes.AllowDeleteAPI {
		delRoute, delRErr := rm.GetDeleteRoute(rm)
		if delRErr == nil {
			r.HandleFunc(delRoute, controller.DeleteHandler).Methods(http.MethodDelete)
		}
	}

	return r
}

func generalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Method implementation not found")
}
