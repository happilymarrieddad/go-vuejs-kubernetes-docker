package v1

import (
	"net/http"

	"github.com/go-xorm/xorm"
)

// Health - health route
func Health(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("V1 Routes are active!"))
	}
}
