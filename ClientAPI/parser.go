package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/tonky/shipping/PortDomainService/domain"
)

// channels in Go are a nice streaming abstraction, allowing to nicely compose with GRPC streaming,
// using mimimal resources on json stream parsing
// io.Reader is another one, making this function pure and testable
// TODO: error handling should be actually implemented
func streamDecodePorts(r io.Reader) (chan domain.Port, chan error) {
	pc := make(chan domain.Port)
	errCh := make(chan error)

	go func() {
		count := 0

		dec := json.NewDecoder(r)

		if _, err := dec.Token(); err != nil {
			log.Println("Error parsing token: ", err)
			errCh <- err
			return
		}

		for dec.More() {
			tok, err := dec.Token()

			if err != nil {
				log.Println("Error parsing key: ", err)
				errCh <- err
				return
			}

			var jp jsPort

			if err := dec.Decode(&jp); err != nil {
				log.Println("Error parsing port: ", err)
				errCh <- err
				return
			}

			key, ok := tok.(string)

			if !ok {
				log.Println("Key doesn't seem to be a string: ", err)
				errCh <- err
				return
			}

			pc <- jsToDomain(key, jp)

			count++
		}

		_, err := dec.Token()

		if err != nil {
			log.Println("Error parsing last token: ", err)
			errCh <- err
			return
		}

		log.Printf("Parsed %d ports, closing chan\n", count)

		close(pc)
		close(errCh)
	}()

	return pc, errCh
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
