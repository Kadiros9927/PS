package auth

import (
	"go/git-ps/http-server/configs"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthHandler_Login(t *testing.T) {
	type fields struct {
		Config *configs.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &AuthHandler{
				Config: tt.fields.Config,
			}
			if got := handler.Login(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthHandler_Register(t *testing.T) {
	type fields struct {
		Config *configs.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &AuthHandler{
				Config: tt.fields.Config,
			}
			if got := handler.Register(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthHandler(t *testing.T) {
	type args struct {
		router *http.ServeMux
		deps   AuthHandlerDeps
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewAuthHandler(tt.args.router, tt.args.deps)
		})
	}
}
