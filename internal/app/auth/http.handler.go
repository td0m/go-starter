package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/td0m/go-starter/pkg/errors"
)

// HTTP handler struct
type HTTP struct {
	svc *Service
}

// NewHTTP attaches router http endpoints
func NewHTTP(r *mux.Router, svc *Service, withJWT mux.MiddlewareFunc) {
	h := HTTP{svc}
	withAuth := r.NewRoute().Subrouter()
	withAuth.Use(withJWT)

	r.HandleFunc("/github", h.authWithGithub).Methods("GET")
	r.HandleFunc("/github/callback", h.authWithGithubCallback).Methods("GET")
	withAuth.HandleFunc("/me", h.me).Methods("POST")
}

func (h *HTTP) me(w http.ResponseWriter, r *http.Request) {
	errors.Write(w, errors.New(http.StatusNotFound, "err"))
}

func (h *HTTP) authWithGithub(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.svc.GithubAuthURL("nostatexx", "http://localhost:8080/auth/github/callback"), http.StatusTemporaryRedirect)
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
