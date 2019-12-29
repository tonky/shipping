package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	domain "github.com/tonky/shipping/PortDomainService/domain"
	pb "github.com/tonky/shipping/PortDomainService/pb"
)

type portsServer struct {
	pb.UnimplementedPortsServer
	storage Storage
}

func (ps *portsServer) Get(ctx context.Context, req *pb.SearchRequest) (*pb.MaybePort, error) {
	fmt.Println("Client request: ", req)

	p, err := ps.storage.Get(req.Name)

	if err != nil {
		return nil, err
	}

	log.Println("Found port in storage: ", p)

	if p == nil {
		return &pb.MaybePort{Found: false, Port: &pb.Port{}}, nil
	}

	port := pb.FromDomain(*p)

	return &pb.MaybePort{Found: true, Port: &port}, nil
}

func (ps *portsServer) Upsert(stream pb.Ports_UpsertServer) error {
	ports := []domain.Port{}

	startTime := time.Now()

	for {
		pbPort, err := stream.Recv()

		if err == io.EOF {
			endTime := time.Now()

			if err := ps.storage.Add(ports); err != nil {
				log.Println("Error adding ports to storage")
				return err
			}

			return stream.SendAndClose(&pb.UpsertSummary{
				Count:       int64(len(ports)),
				ElapsedTime: endTime.Sub(startTime).Milliseconds(),
			})
		}

		if err != nil {
			return err
		}

		port := pb.ToDomain(*pbPort)

		ports = append(ports, port)
	}
}
