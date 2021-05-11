package main

import (
	"log"
	"net/http"

	"github.com/onprem/cordi/pkg/api"
)

func main() {
	router := http.NewServeMux()

	_ = api.New(router, "foo", "secret")

	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Println("server listen error: ", err.Error())
	}
}
