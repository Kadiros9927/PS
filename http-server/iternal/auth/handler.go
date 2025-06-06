package auth

import (
	"fmt"
	"go/git-ps/http-server/configs"
	"go/git-ps/http-server/pkg/req"
	"go/git-ps/http-server/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct { //
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) { // NewAuthHandler
	// initializes the auth handler and registers routes
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *AuthHandler) Login() http.HandlerFunc { // Login handler function
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r) // Decode and validate request body
		if err != nil {
			return // If there's an error, it has already been handled in HandleBody
		}
		fmt.Println(body)
		data := LoginPayload{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc { // Register handler function
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r) // Decode and validate request body
		if err != nil {
			return // If there's an error, it has already been handled in HandleBody
		}
		fmt.Println(body)
	}
}
