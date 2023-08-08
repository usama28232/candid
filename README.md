# potential-candid-go
A minimal application to replicate basic web api crud opertaions.

Purpose of this project is to provide most a commonly used base-line application to assist development.

The goal is to eliminate repetitive tasks and focus on writing the business aspect.

## Project Structure & Explanation

The project structure consists of a mixture of MVC & Service based architecture.


### Controllers

This (package: `controllers`) will contain the default handlers for `GET`, `POST`, `DELETE`, etc requests. It exposes the following interface in `controller_base.go`

In other words: To make the most of it, a controller must follow this pattern.

```
type controllerBase interface {
	// GET: "<context>/<base_route>"
	GetAllHandler(http.ResponseWriter, *http.Request)

	// POST: "<context>/<base_route>"
	CreateHandler(http.ResponseWriter, *http.Request)

	// GET: "<context>/<base_route>/<value>"
	GetHandler(http.ResponseWriter, *http.Request)

	// DELETE: "<context>/<base_route>/<value>"
	DeleteHandler(http.ResponseWriter, *http.Request)

	GetRouteModel() routes.RouteConfig
}
```

`GetRouteModel` returns an 'derived' instance of controller's route configuration *- explained below*

For example, a simple controller definition `HelloController` would be:

```
...
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
...
```

This is the engagement point of service layer code.


### Route-Model

RouteModel (package: `routes`) is the basic skeleton of what a 'derived' controller is capable of, the above was all about the landing points. This area deals with the mapping between routes and landing points.

Default definition & Implementation from `RouteModel.go`

```
// Holds allowed endpoint configuration
type AllowedRoutes struct {
	AllowListAPI bool
	AllowAddAPI bool
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
```

From this, you will be able to add a derived implementation by doing just this:

```
type HelloRouteModel struct {
	routes.RouteConfigImpl
}

func (r *HelloRouteModel) Init() routes.AllowedRoutes {
	return routes.AllowedRoutes{AllowListAPI: true, AllowDeleteAPI: true}
}

func (r *HelloRouteModel) GetBaseRoute() (string, error) {
	return "/hello", nil
}
```

The above definition will tell the framework to map base route `/hello` to `HelloController` *- see above example*

The `Init` method returns the allowed endpoints configuration in the following format.

```
type AllowedRoutes struct {
	AllowListAPI bool
	AllowAddAPI bool
	AllowDetailAPI bool
	AllowDeleteAPI bool
}
```

This is then evaluated during registration process of the framework

```
...
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
...
```

Because of this architecture, router code is reduced to minimal

```
func RegisterRoutes() *mux.Router {

	mux := mux.NewRouter()
	mux.Use(authMiddleware)
	helloCont := &HelloController{}
	userCont := &UserController{}

	mux = register(userCont, mux)
	mux = register(helloCont, mux)

	mux.StrictSlash(false)
	return mux
}
```

So to add an endpoint, all you have to do is to add it's derived route-model and controller with required landings.

### Service

This is the layer where business logic is supposed to live.

Here is an example of `HelloService`

```
func SayHello() string {
	return "Hello World!"
}
```
Which will be called by derived the controller.

This layer is independent so it can cover all aspects because in some cases there are composite entities that do not require exposed endpoints but are crucial part of a bigger operation.

For advanced example, Follow `UserController`

...

## Conclusion

This framework makes development easier and robust.
You can clone and start writing business right away!

Furthermore, I have plans to add `GORM` and provide basic crud functionality at service layer.



### Feel free to edit/expand/explore this repository

For feedback and queries, reach me on LinkedIn at [here](https://www.linkedin.com/in/usama28232/?original_referer=)