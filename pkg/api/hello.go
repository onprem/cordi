package api

import (
	"fmt"
	"net/http"
)

func (a *Api) sayHello(w http.ResponseWriter, r *http.Request) {
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
