#!/bin/sh
go build -o app && env $(cat .env_local_pg | xargs) ./app
