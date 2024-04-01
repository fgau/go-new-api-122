package controllers

import (
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("User ID: " + userID))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	world := r.PathValue("world")
	// data := SomeStruct{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"hello": world,
	})
}
