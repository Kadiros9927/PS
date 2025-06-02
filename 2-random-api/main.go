package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

type MyMux struct{}

func (mux *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write([]byte(fmt.Sprintf("%d", rand.IntN(7))))

	}

}
func main() {
	// Custom Mux
	router := &MyMux{}
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}

	server.ListenAndServe()
}
