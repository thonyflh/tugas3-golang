package middleware

import "net/http"

var (
	USERNAME = "anthony"
	PASSWORD = "12345"
)

func MiddlewareBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error un authorized"))
			return
		}

		isValid := (username == USERNAME && password == PASSWORD)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error un authorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
