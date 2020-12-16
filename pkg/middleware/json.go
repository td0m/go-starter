package middleware

import "net/http"

// ContentTypeJSON specifies that the response will be of type JSON using a Content-Type header
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
