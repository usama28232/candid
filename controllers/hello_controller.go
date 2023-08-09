package controllers

import (
	"net/http"

	"github.com/usama28232/candid/hello"
	"github.com/usama28232/candid/routes"
	"github.com/usama28232/candid/shared"
)

type HelloController struct {
	myControllerBase
}

func (c *HelloController) GetAllHandler(w http.ResponseWriter, request *http.Request) {
	shared.EncodeResponse(hello.SayHello(), w)
}

func (cb *HelloController) GetHandler(w http.ResponseWriter, r *http.Request) {
	shared.EncodeResponse(hello.SayHello(), w)
}

func (c *HelloController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	shared.EncodeResponse(hello.SayDelete(), w)
}

func (u *HelloController) GetCustomLanding(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte(hello.SayHello()))
}

func (u *HelloController) GetRouteModel() routes.RouteConfig {
	return &hello.HelloRouteModel{}
}
