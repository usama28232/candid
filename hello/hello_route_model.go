package hello

import "github.com/usama28232/candid/routes"

type HelloRouteModel struct {
	routes.RouteConfigImpl
}

func (r *HelloRouteModel) Init() routes.AllowedRoutes {
	return routes.AllowedRoutes{AllowListAPI: true, AllowDeleteAPI: true, AllowDetailAPI: true}
}

func (r *HelloRouteModel) GetBaseRoute() (string, error) {
	return "/hello", nil
}
