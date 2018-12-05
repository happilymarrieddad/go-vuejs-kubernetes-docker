package routes

import "net/http"

// Routes - routes
type Routes []Route

// Route - route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// SubRoutePackage - sub route package
type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
