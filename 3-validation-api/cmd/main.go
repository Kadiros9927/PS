package main

import (
	"go/git-ps/3-validation-api/config"
	"go/git-ps/3-validation-api/iternal/verify"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	cfg := config.LoadConfig()
	router.HandleFunc("/send", verify.NewVerifyHandler(cfg))
	router.HandleFunc("/verify/", verify.VerifyHandler())

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
