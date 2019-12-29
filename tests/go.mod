module github.com/tonky/shipping/tests

go 1.13

require (
	github.com/tonky/shipping/ClientAPI v0.0.0
	github.com/tonky/shipping/PortDomainService v0.0.0
)

replace github.com/tonky/shipping/ClientAPI => ../ClientAPI

replace github.com/tonky/shipping/PortDomainService => ../PortDomainService
