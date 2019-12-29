package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tonky/shipping/ClientAPI/portService"
	"github.com/tonky/shipping/PortDomainService/domain"
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

type jsPort struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

func jsToDomain(key string, jp jsPort) domain.Port {
	p := domain.Port{
		Key:          key,
		Name:         jp.Name,
		City:         jp.City,
		Country:      jp.Country,
		Aliases:      jp.Alias,
		Regions:      jp.Regions,
		Province:     jp.Province,
		LocationName: jp.Timezone,
		Unlocs:       jp.Unlocs,
		Code:         jp.Code,
	}

	if len(jp.Coordinates) == 2 {
		p.Latitude = jp.Coordinates[0]
		p.Longitude = jp.Coordinates[1]
	} else {
		log.Println("Bad coordinates: ", key, jp)
	}

	return p
}
