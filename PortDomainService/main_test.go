package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	domain "github.com/tonky/shipping/PortDomainService/domain"
	pb "github.com/tonky/shipping/PortDomainService/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type testStorage struct {
	ports map[string]domain.Port
}

func (ts testStorage) Add(ports []domain.Port) error {
	log.Println("Adding ports: ", ports)

	for _, p := range ports {
		ts.ports[p.Key] = p
	}

	log.Println("added ports: ", ts.ports)

	return nil
}

func (ts testStorage) Get(key string) (*domain.Port, error) {
	log.Printf("testStorage Get(%s) from %v\n", key, ts.ports)

	p := ts.ports[key]

	return &p, nil
}

func init() {
	lis = bufconn.Listen(bufSize)

	s := grpc.NewServer()

	ps := map[string]domain.Port{}

	ts := testStorage{ps}

	pb.RegisterPortsServer(s, &portsServer{storage: ts})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestService(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()

	client := pb.NewPortsClient(conn)

	stream, err := client.Upsert(ctx)

	if err != nil {
		t.Fatalf("Upsert stream failed: %v", err)
	}

	ports := []pb.Port{
		{Key: "UAIEV", Name: "Kiev", City: "Kiev", Country: "Ukraine", Code: "54112"},
		{Key: "ANCUR", Name: "Curacao", City: "Curacao", Country: "Netherlands Antilles"},
		{Key: "GBBRS", Name: "Bristol", City: "Bristol", Code: "41211", Country: "United Kingdom"},
	}

	for _, port := range ports {
		if err := stream.Send(&port); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, port, err)
		}
	}

	reply, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}

	log.Println("Got RPC reply: ", reply)

	if reply.Count != int64(len(ports)) {
		t.Fatal("Expected 1 acknowledge")
	}

	maybePort, err := client.Get(ctx, &pb.SearchRequest{Name: "GBBRS"})

	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}

	port := maybePort.Port

	if !maybePort.Found || port.Key != "GBBRS" || port.Code != "41211" {
		t.Fatalf("Got: %+v\n", port)
	}
}
