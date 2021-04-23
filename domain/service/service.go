package service

import "github.com/victorrub/dns-reset/infra/network"

// Service holds the domain service repositories
type Service struct {
	net *network.Communicator
}

// New returns a new domain Service instance
func New(net *network.Communicator) (*Service, error) {
	svc := new(Service)

	svc.net = net

	return svc, nil
}
