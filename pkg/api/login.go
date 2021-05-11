package api

import (
	"encoding/json"
	"net/http"
)

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
