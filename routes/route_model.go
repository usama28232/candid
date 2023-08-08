package routes

import "errors"

type AllowedRoutes struct {
	AllowListAPI   bool
	AllowAddAPI    bool
	AllowDetailAPI bool
	AllowDeleteAPI bool
}

type RouteConfig interface {
	Init() AllowedRoutes
	GetBaseRoute() (string, error)
	GetListRoute(RouteConfig) (string, error)
	GetAddRoute(RouteConfig) (string, error)
	GetInfoRoute(RouteConfig) (string, error)
	GetDeleteRoute(RouteConfig) (string, error)
}

type RouteConfigImpl struct {
}

func (r *RouteConfigImpl) Init() AllowedRoutes {
	return AllowedRoutes{AllowListAPI: false, AllowAddAPI: false, AllowDetailAPI: false, AllowDeleteAPI: false}
}

func (r *RouteConfigImpl) GetBaseRoute() (string, error) {
	return r.RouteErrHandler()
}

func (r *RouteConfigImpl) GetListRoute(c RouteConfig) (string, error) {
	return c.GetBaseRoute()
}

func (r *RouteConfigImpl) GetAddRoute(c RouteConfig) (string, error) {
	return c.GetBaseRoute()
}

func (r *RouteConfigImpl) GetInfoRoute(c RouteConfig) (string, error) {
	v, err := c.GetBaseRoute()
	return v + "/{value:[A-Za-z0-9]+}", err
}

func (r *RouteConfigImpl) GetDeleteRoute(c RouteConfig) (string, error) {
	v, err := c.GetBaseRoute()
	return v + "/{value:[A-Za-z0-9]+}", err
}

func (r *RouteConfigImpl) RouteErrHandler() (string, error) {
	return "", errors.New("error: route config not defined")
}
