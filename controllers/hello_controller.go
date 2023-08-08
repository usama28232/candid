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

func (c *HelloController) GetAllHandler(writer http.ResponseWriter, request *http.Request) {
	shared.EncodeResponse(hello.SayHello(), writer)
}

func (c *HelloController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	shared.EncodeResponse("Success", w)
}

func (u *HelloController) GetRouteModel() routes.RouteConfig {
	return &hello.HelloRouteModel{}
}
