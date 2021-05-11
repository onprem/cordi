package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Api struct {
	DB map[string]string
}

func New(router *http.ServeMux, username string, password string) *Api {
	a := &Api{
		DB: map[string]string{
			"user":     username,
			"password": password,
		},
	}
	a.registerRoutes(router)
	return a
}

func (a *Api) sayPong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

type resp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
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
