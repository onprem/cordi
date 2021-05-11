package api

import "net/http"

func (a *Api) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", a.sayPong)
	mux.HandleFunc("/hello", a.sayHello)
	mux.HandleFunc("/login", a.handleLogin)
}
