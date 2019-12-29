FROM golang:alpine
WORKDIR /src
ADD . .
ENTRYPOINT CGO_ENABLED=0 go test -mod=vendor
