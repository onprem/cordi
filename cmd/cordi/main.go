package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		returnError(w, "malformed request")
		return
	}

	if input.Email != "foo" || input.Password != "bar" {
		w.WriteHeader(http.StatusUnauthorized)
		returnError(w, "invalid credentials")
		return
	}

	returnSuccess(w, "valid credentials")
}

func returnSuccess(w io.Writer, msg string) {
	returnResp(w, "success", msg)
}

func returnError(w io.Writer, msg string) {
	returnResp(w, "error", msg)
}

func returnResp(w io.Writer, status string, msg string) {
	if err := json.NewEncoder(w).Encode(resp{
		Status: status,
		Msg:    msg,
	}); err != nil {
		log.Println("error encoding response: ", err.Error())
	}
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/ping", sayPong)
	router.HandleFunc("/hello", sayHello)
	router.HandleFunc("/login", handleLogin)

	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Println("server listen error: ", err.Error())
	}
}
