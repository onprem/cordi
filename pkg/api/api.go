package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func sayPong(w http.ResponseWriter, r *http.Request) {
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
