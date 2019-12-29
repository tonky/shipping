module github.com/tonky/shipping/ClientAPI

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/tonky/shipping/PortDomainService v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.26.0
)

replace github.com/tonky/shipping/PortDomainService => ../PortDomainService
