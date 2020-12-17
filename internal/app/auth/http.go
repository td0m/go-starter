package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/td0m/go-starter/pkg/errors"
	"github.com/td0m/go-starter/pkg/jwt"
)

// HTTP handler struct
type HTTP struct {
	svc Service
}

// NewHTTP attaches router http endpoints
func NewHTTP(r *mux.Router, svc Service, auth mux.MiddlewareFunc) {
	h := HTTP{svc}
	withAuth := r.NewRoute().Subrouter()
	withAuth.Use(auth)

	r.HandleFunc("/github", h.authWithGithub).Methods("GET")
	r.HandleFunc("/github/callback", h.authWithGithubCallback).Methods("GET")
	withAuth.HandleFunc("/me", h.me).Methods("GET")
}

func (h *HTTP) me(w http.ResponseWriter, r *http.Request) {
	user := jwt.FromContext(r.Context())
	json.NewEncoder(w).Encode(user)
}

func (h *HTTP) authWithGithub(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.svc.GithubAuthURL("nostatexx"), http.StatusTemporaryRedirect)
}

func (h *HTTP) authWithGithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := h.svc.GithubCodeToToken(code)
	if err != nil {
		errors.Write(w, err)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
