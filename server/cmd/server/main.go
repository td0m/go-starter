package main

import (
	"net/http"

	"github.com/d0minikt/go-starter/server/pkg/add"
	"github.com/d0minikt/go-starter/server/pkg/get"
	"github.com/d0minikt/go-starter/server/pkg/handlers/rest"
	"github.com/d0minikt/go-starter/server/pkg/storage/postgres"
)

func main() {
	//restService := rest.NewService(getService)
	//http.ListenAndServe(":8080", restService.Handler())
	repo := postgres.New()
	g := get.NewService(repo)
	a := add.NewService(repo)
	rest := rest.New(g, a)
	http.ListenAndServe(":8080", rest.Handler())
}
