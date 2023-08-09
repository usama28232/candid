package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/usama28232/candid/logging"
	"github.com/usama28232/candid/routes"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type AppRequestLog struct {
	Method   string
	Url      string
	Data     string
	Agent    string
	Duration int64
}

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
	generalError(w, r)
}

func (cb *myControllerBase) CreateHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w, r)
}

func (cb *myControllerBase) GetHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w, r)
}

func (cb *myControllerBase) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	generalError(w, r)
}

func (cb *myControllerBase) GetRouteModel() (routes.RouteConfig, error) {
	return nil, errors.New("route model not configured correctly")
}

func registerCustom(path string, f http.HandlerFunc, r *mux.Router) *mux.Router {
	logger := logging.GetLogger()
	logger.Debugw("- Adding custom route", "v", path)
	r.HandleFunc(path, f)
	return r
}

func register(controller controllerBase, r *mux.Router) *mux.Router {
	logger := logging.GetLogger()
	rm := controller.GetRouteModel()
	availableRoutes := rm.Init()
	baseRoute, baseRErr := rm.GetBaseRoute()
	logger.Infow("** Route Config **", "Route", baseRoute, zap.Any("Available Routes", availableRoutes))

	if baseRErr == nil && availableRoutes.AllowListAPI {
		r.HandleFunc(baseRoute, controller.GetAllHandler).Methods(http.MethodGet)
		logger.Debugw("- Adding Route", "v", baseRoute, "m", http.MethodGet)
	}

	if availableRoutes.AllowAddAPI {
		addRoute, addRErr := rm.GetAddRoute(rm)
		if addRErr == nil {
			r.HandleFunc(addRoute, controller.CreateHandler).Methods(http.MethodPost)
			logger.Debugw("- Adding Route", "v", addRoute, "m", http.MethodPost)
		}
	}

	if availableRoutes.AllowDetailAPI {
		getInfoRoute, getInfoRErr := rm.GetInfoRoute(rm)
		if getInfoRErr == nil {
			r.HandleFunc(getInfoRoute, controller.GetHandler).Methods(http.MethodGet)
			logger.Debugw("- Adding Route", "v", getInfoRoute, "m", http.MethodGet)
		}
	}

	if availableRoutes.AllowDeleteAPI {
		delRoute, delRErr := rm.GetDeleteRoute(rm)
		if delRErr == nil {
			r.HandleFunc(delRoute, controller.DeleteHandler).Methods(http.MethodDelete)
			logger.Debugw("- Adding Route", "v", delRoute, "m", http.MethodDelete)
		}
	}

	return r
}

func generalError(w http.ResponseWriter, r *http.Request) {

	meta := AppRequestLog{}
	meta.Method = r.Method
	meta.Url = r.URL.Path
	meta.Agent = r.UserAgent()

	logging.GetAccessLogger().Infow("HttpRequest", zap.Any("v", meta))

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Method implementation not found")
}
