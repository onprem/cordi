package main

import (
	"net/http"
)

func sayPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", sayPong)

	http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
}
