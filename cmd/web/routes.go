package main

import (
	"github.com/bmizerany/pat"
	"github.com/rostis232/awesomeProject/pkg/config"
	"github.com/rostis232/awesomeProject/pkg/handlers"
	"net/http"
)

func routes(a *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
