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
