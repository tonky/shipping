FROM golang as builder
ADD . /src
RUN cd /src && CGO_ENABLED=0 go build -mod=vendor -o api

FROM alpine
WORKDIR /app
COPY --from=builder /src/api /app/
ENTRYPOINT ./api
