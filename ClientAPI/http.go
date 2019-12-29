package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tonky/shipping/ClientAPI/portService"
)

type api struct {
	portService portService.Api
}

func (ah api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("request form: ", r.URL.Query())

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "query 'name' can't be empty", http.StatusBadRequest)
		return
	}

	port, err := ah.portService.Get(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if port == nil {
		http.Error(w, "Port not found", http.StatusNotFound)
		return
	}

	jp, err := json.Marshal(port)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jp)
}
