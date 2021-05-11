package main

import (
	"log"
	"net/http"

	"github.com/onprem/cordi/pkg/api"
)

func main() {
	router := http.NewServeMux()
	api.RegisterRoutes(router)

	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Println("server listen error: ", err.Error())
	}
}
