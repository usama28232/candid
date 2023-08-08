package hello

import "authexample/routes"

type HelloRouteModel struct {
	routes.RouteConfigImpl
}

func (r *HelloRouteModel) Init() routes.AllowedRoutes {
	return routes.AllowedRoutes{AllowListAPI: true, AllowDeleteAPI: true}
}

func (r *HelloRouteModel) GetBaseRoute() (string, error) {
	return "/hello", nil
}
