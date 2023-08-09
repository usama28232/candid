package controllers

import (
	"authexample/hello"
	"authexample/routes"
	"authexample/shared"
	"net/http"
)

type HelloController struct {
	myControllerBase
}

func (c *HelloController) GetAllHandler(w http.ResponseWriter, request *http.Request) {
	shared.EncodeResponse(hello.SayHello(), w)
}

func (c *HelloController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	shared.EncodeResponse(hello.SayDelete(), w)
}

func (u *HelloController) GetRouteModel() routes.RouteConfig {
	return &hello.HelloRouteModel{}
}
