package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/td0m/go-starter/internal/app/shortener"
)

// HTTP handler
type HTTP struct {
	svc shortener.Service
}

// New creates a new app http handler
func New(r *mux.Router, svc shortener.Service) {
	h := HTTP{svc}

	r.HandleFunc("/{id}", h.get).Methods("GET")
	r.HandleFunc("", h.create).Methods("POST")
}

type helloResponse struct {
	Time time.Time `json:"time"`
	To   string    `json:"to"`
}

func (h *HTTP) get(w http.ResponseWriter, r *http.Request) {
	l, err := h.svc.Get(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, l.URL, http.StatusTemporaryRedirect)
}

func (h *HTTP) create(w http.ResponseWriter, r *http.Request) {
	body := new(struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	})
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	l, err := h.svc.Create(body.ID, body.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(l)
}
