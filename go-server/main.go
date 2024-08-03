package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{
		"ok": true,
	})
}
