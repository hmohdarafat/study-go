package main

import (
	"fmt"
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
			case "/":
				w.Write([]byte("index page"))
				return
			case "/users":
				w.Write([]byte("users page"))
				return
			default:
				w.Write([]byte("404 page"))
				return
		}
	default:
		w.Write([]byte("404 page"))
		return
	}
}

func main() {
	fmt.Println("HTTP Server Started...")
	api := &api{
		addr: ":8080",
	}
	srv := &http.Server{
		Addr: api.addr,
		Handler: api,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}