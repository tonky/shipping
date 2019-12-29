package main

import domain "github.com/tonky/shipping/PortDomainService/domain"

type Storage interface {
	Add([]domain.Port) error
	Get(string) (*domain.Port, error)
}
