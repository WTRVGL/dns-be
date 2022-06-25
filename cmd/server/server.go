package main

import (
	"encoding/json"
	dns_be "github.com/WTRVGL/dns-be"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", CheckDomain)
	http.ListenAndServe("localhost:8080", mux)
}

func CheckDomain(w http.ResponseWriter, req *http.Request) {
	runes := []rune(req.URL.Path)
	name := string(runes[1:])
	domain, err := dns_be.NewDomain(name)

	if err != nil {
		w.Write([]byte(err.Error()))
		http.NotFound(w, req)
		return
	}

	domain.CheckAvailability()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(domain)
}
