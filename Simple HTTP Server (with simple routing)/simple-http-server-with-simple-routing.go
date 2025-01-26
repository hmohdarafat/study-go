package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
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
	s := &server{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}