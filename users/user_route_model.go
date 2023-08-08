package users

import "authexample/routes"

type UserRouteModel struct {
	routes.RouteConfigImpl
}

func (r *UserRouteModel) Init() routes.AllowedRoutes {
	return routes.AllowedRoutes{AllowListAPI: true, AllowAddAPI: true, AllowDetailAPI: true, AllowDeleteAPI: true}
}

func (r *UserRouteModel) GetBaseRoute() (string, error) {
	return "/users", nil
}
