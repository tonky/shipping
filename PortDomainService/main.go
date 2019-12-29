package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	pb "github.com/tonky/shipping/PortDomainService/pb"
	"github.com/tonky/shipping/PortDomainService/postgres"
	"google.golang.org/grpc"
)

func main() {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbname := os.Getenv("DBNAME")

	db, err := postgres.New(user, password, host, port, dbname)

	if err != nil {
		log.Fatalln("Can't connect to storage: ", err)
	}

	addr := os.Getenv("DOMAIN_SERVICE_ADDR")

	if addr == "" {
		addr = ":1234"
	}

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPortsServer(grpcServer, &portsServer{storage: db})

	fmt.Println("Starting RPC server on ", addr)
	grpcServer.Serve(lis)
}
