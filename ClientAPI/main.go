package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tonky/shipping/ClientAPI/portService"
)

func main() {
	hp := os.Getenv("HTTP_PORT")
	psAddr := os.Getenv("DOMAIN_SERVICE_ADDR")
	fileToLoad := os.Getenv("LOAD_FILE")

	if hp == "" {
		hp = "8080"
	}

	if psAddr == "" {
		log.Fatal("DOMAIN_SERVICE_ADDR is not set")
	}

	ps, err := portService.New(psAddr)

	if err != nil {
		log.Println("Can't get portService: ", err)
	}

	if fileToLoad != "" {
		file, err := os.Open(fileToLoad)

		if err != nil {
			log.Fatal("Can't open file: ", err)
		}

		// TODO: decide on hadling errors from error chan
		pc, _ := streamDecodePorts(file)

		go ps.Send(pc)
	}

	http.Handle("/", api{ps})

	log.Println("Starting http service on port ", hp)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", hp), nil))
}
