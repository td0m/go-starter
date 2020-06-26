package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/d0minikt/go-starter/server/pkg/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type handler struct {
	userService domain.UserService
}

func New(userService domain.UserService) handler {
	return handler{userService}
}

func (h *handler) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", h.list).Methods("GET")
	r.HandleFunc("/auth/email", h.authByEmail).Methods("POST")
	r.HandleFunc("/api/users/{email}", h.withClaims(h.getUser)).Methods("GET")
	return r
}

type emailAuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) authByEmail(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var body emailAuthBody
	if err := d.Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, token, err := h.userService.Login(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (h *handler) getUser(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]

	claims := r.Context().Value("claims").(jwt.MapClaims)
	if claims["sub"] != email {
		http.Error(w, "Emails didnt match", http.StatusUnauthorized)
		return
	}

	user, err := h.userService.GetUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "oof")
}
