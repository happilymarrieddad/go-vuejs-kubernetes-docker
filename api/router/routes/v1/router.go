package v1

import (
	"net/http"

	"github.com/go-xorm/xorm"
	Jwt "github.com/learning/project/api/jwt"
	Routes "github.com/learning/project/api/router/routes"
)

// Middleware - this is the middleware for v1 routes
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")

		if _, err := Jwt.ParseToken(token); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1 routes
func GetRoutes(db *xorm.Engine) (SubRoute map[string]Routes.SubRoutePackage) {
	SubRoute = map[string]Routes.SubRoutePackage{
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name:        "V1HealthRoute",
					Method:      "GET",
					Pattern:     "/health",
					HandlerFunc: Health(db),
				},
			},
			Middleware: Middleware,
		},
	}

	return SubRoute
}
