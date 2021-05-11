package api

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", sayPong)
	mux.HandleFunc("/hello", sayHello)
	mux.HandleFunc("/login", handleLogin)
}
