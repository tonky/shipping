#!/bin/sh
# LOAD_FILE="`pwd`/../ports.json" ./run.sh
go build -o api && DOMAIN_SERVICE_ADDR=localhost:1234 LOAD_FILE="`pwd`/../sampleData/ports.json" ./api
