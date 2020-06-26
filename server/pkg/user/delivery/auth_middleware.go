package delivery

import (
	"context"
	"net/http"
	"strings"
)

func (h *handler) withClaims(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if !strings.Contains(tokenStr, " ") {
			http.Error(w, "no authorization header", http.StatusBadRequest)
			return
		}
		tokenStr = strings.Split(tokenStr, " ")[1]
		claims, err := h.userService.GetTokenClaims(tokenStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		f(w, r.WithContext(ctx))
	}
}
