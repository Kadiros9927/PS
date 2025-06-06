package main

import (
	"go/git-ps/http-server/configs"
	"go/git-ps/http-server/iternal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
