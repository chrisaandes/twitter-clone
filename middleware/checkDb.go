package middleware

import (
	"net/http"

	"github.com/chrisaandes/twitter-clone/db"
)

/*CheckDb ...*/
func CheckDb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "db missed", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
