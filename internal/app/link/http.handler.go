package link

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
func NewHTTP(r *mux.Router, svc *Service) {
	h := HTTP{svc}

	r.HandleFunc("/{id}", h.get).Methods("GET")
	r.HandleFunc("", h.create).Methods("POST")
	r.HandleFunc("/{id}", h.put).Methods("PUT")
}

func (h *HTTP) get(w http.ResponseWriter, r *http.Request) {
	l, err := h.svc.Get(mux.Vars(r)["id"])
	if err != nil {
		errors.Write(w, err)
		return
	}
	http.Redirect(w, r, l.Url, http.StatusTemporaryRedirect)
}

type putBody struct {
	URL string `json:"url"`
}

func (h *HTTP) put(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	body := new(putBody)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		errors.WithCode(http.StatusBadRequest, err).Write(w)
		return
	}
	l, err := h.svc.Create(id, body.URL)
	if err != nil {
		errors.Write(w, err)
		return
	}
	json.NewEncoder(w).Encode(l)
}

type createBody struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (h *HTTP) create(w http.ResponseWriter, r *http.Request) {
	body := new(createBody)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		errors.WithCode(400, err).Write(w)
		return
	}
	l, err := h.svc.Create(body.ID, body.URL)
	if err != nil {
		errors.Write(w, err)
		return
	}
	json.NewEncoder(w).Encode(l)
}
