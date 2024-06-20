package main

import (
	"net/http"

	"github.com/pisarivskyi/goapi"
)

func main() {
	server := goapi.NewServer("127.0.0.1:1234")

	server.Router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})

	server.Router.Get("/test/test/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test/test/test"))
	})

	server.Router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server.Router.Get("/hello/world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello/world"))
	})

	server.Serve()
}
