package rest

import (
	"encoding/json"
	"net/http"

	"github.com/d0minikt/go-starter/server/pkg/add"
	"github.com/d0minikt/go-starter/server/pkg/get"
	"github.com/gorilla/mux"
)

type Rest struct {
	get get.Service
	add add.Service
}

func New(g get.Service, a add.Service) Rest {
	return Rest{g, a}
}

func (re *Rest) Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/links", re.getLinks).Methods("GET")
	r.HandleFunc("/api/links", re.createLink).Methods("POST")
	r.HandleFunc("/{alias}", re.putLink).Methods("PUT")
	r.HandleFunc("/{alias}", re.getLink).Methods("GET")
	return r
}
func (self *Rest) putLink(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	d := json.NewDecoder(r.Body)
	var body add.Link
	if err := d.Decode(&body); err != nil {
		//HANDLE ERROR
	}
	body.Alias = alias
	self.add.AddLink(body)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (self *Rest) getLink(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	l, _ := self.get.GetLinkByAlias(alias)
	http.Redirect(w, r, l.URL, 301)
}

func (self *Rest) createLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	d := json.NewDecoder(r.Body)
	var body add.Link
	if err := d.Decode(&body); err != nil {
		//HANDLE ERROR
	}
	self.add.AddLink(body)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (re *Rest) getLinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(re.get.GetLinks())
}
