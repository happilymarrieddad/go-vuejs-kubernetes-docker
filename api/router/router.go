package router

import (
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"

	Routes "github.com/learning/project/api/router/routes"
	V1Routes "github.com/learning/project/api/router/routes/v1"
)

const (
	staticDir = "/static/"
)

// RouteHandler - the handler for go api routes
type RouteHandler struct {
	Router *mux.Router
}

// AttachSubRouterWithMiddleware - allows you to attach
// 		subroutes to router
func (r *RouteHandler) AttachSubRouterWithMiddleware(
	path string,
	subroutes Routes.Routes,
	Middleware mux.MiddlewareFunc,
) (SubRouter *mux.Router) {

	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(Middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

// NewRouter - create a new router
func NewRouter(db *xorm.Engine) *RouteHandler {
	var router RouteHandler

	router.Router = mux.NewRouter().StrictSlash(true)

	router.Router.PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	router.Router.Use(Routes.Middleware)

	routes := Routes.GetRoutes(db)

	for _, route := range routes {
		router.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	v1SubRoutes := V1Routes.GetRoutes(db)

	for name, pack := range v1SubRoutes {
		router.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}

	return &router
}
