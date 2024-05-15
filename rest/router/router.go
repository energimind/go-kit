package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// route is an interface that represents a collection of routes,
// that can be registered to a router.
type route interface {
	RegisterRoutes(root gin.IRouter)
}

// routeFunc is a function that registers routes to a router.
type routeFunc func(root gin.IRouter)

// Router is the router for the REST API.
type Router struct {
	rtr *gin.Engine
}

// New creates a new Router.
func New(middleware ...gin.HandlerFunc) *Router {
	gin.SetMode(gin.ReleaseMode)

	rtr := gin.New()

	rtr.Use(middleware...)

	return &Router{rtr: rtr}
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.rtr.ServeHTTP(w, req)
}

// RegisterRoutes registers the routes to the router.
func (r *Router) RegisterRoutes(routes route) {
	routes.RegisterRoutes(r.rtr)
}

// RegisterRoutesFunc registers the routes to the router using a function.
func (r *Router) RegisterRoutesFunc(fn routeFunc) {
	fn(r.rtr)
}

// GetRoutes returns the routes of the router.
func (r *Router) GetRoutes() []RouteInfo {
	routes := r.rtr.Routes()
	infos := make([]RouteInfo, 0, len(routes))

	for _, rt := range routes {
		infos = append(infos, RouteInfo{
			Method: rt.Method,
			Path:   rt.Path,
		})
	}

	return infos
}
