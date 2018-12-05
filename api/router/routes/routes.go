package routes

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
)

// Middleware - this is the main middleware for the application
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Main middleware")
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - basic routes
func GetRoutes(db *xorm.Engine) Routes {
	return Routes{
		Route{
			Name:        "HealthCheck",
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: Health(db),
		},
		Route{
			Name:        "Login",
			Method:      "POST",
			Pattern:     "/auth/login",
			HandlerFunc: Login(db),
		},
		Route{
			Name:        "Chec",
			Method:      "POST",
			Pattern:     "/auth/check",
			HandlerFunc: Check(db),
		},
	}
}
