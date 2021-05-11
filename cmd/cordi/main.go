package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sayPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("malformed request"))
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name query param is required"))
		return
	}

	fmt.Fprintf(w, "Hello %s", name)
}

type resp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp{
			Status: "error",
			Msg:    "malformed request",
		})
		return
	}

	if input.Email != "foo" || input.Password != "bar" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp{
			Status: "error",
			Msg:    "invalid credentials",
		})
		return
	}

	json.NewEncoder(w).Encode(resp{
		Status: "success",
		Msg:    "valid credentials",
	})
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/ping", sayPong)
	router.HandleFunc("/hello", sayHello)
	router.HandleFunc("/login", handleLogin)

	http.ListenAndServe("0.0.0.0:8080", router)
}
