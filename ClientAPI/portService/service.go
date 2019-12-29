package portService

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	"github.com/tonky/shipping/PortDomainService/domain"
	pb "github.com/tonky/shipping/PortDomainService/pb"
	"google.golang.org/grpc"
)

type Api interface {
	Get(string) (*domain.Port, error)
	Send(chan domain.Port) error
}

type pService struct {
	client pb.PortsClient
}

func (ps pService) Get(name string) (*domain.Port, error) {
	mp, err := ps.client.Get(context.Background(), &pb.SearchRequest{Name: name})

	if err != nil {
		return nil, err
	}

	if !mp.Found {
		log.Println("Port not found: ", mp)

		return nil, nil
	}

	pbPort := mp.Port

	log.Println("Got port: ", pbPort)

	port := pb.ToDomain(*pbPort)

	return &port, nil
}

func (ps pService) Send(ports chan domain.Port) error {
	var stream pb.Ports_UpsertClient
	retries := 5

	for i := 1; i <= retries; i++ {
		var err error

		stream, err = ps.client.Upsert(context.Background())

		if err == nil {
			break
		} else if i == retries {
			return err
		} else {
			log.Println("Error getting client.Upsert stream, will try after some sleep")
		}

		time.Sleep(700 * time.Millisecond)
	}

	if stream == nil {
		return errors.New("client.Upsert failed")
	}

	for port := range ports {
		pbPort := pb.FromDomain(port)

		if err := stream.Send(&pbPort); err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("%v.Send(%v) = %v", stream, pbPort, err)
		}
	}

	reply, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		return err
	}

	log.Printf("Sent summary: %v", reply)

	return nil
}

func New(addr string) (*pService, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	log.Println("Dialed ", addr)

	return &pService{pb.NewPortsClient(conn)}, nil
}
