package service

import (
	"fmt"

	"github.com/victorrub/dns-reset/domain"
	"github.com/victorrub/dns-reset/domain/contract"
	"github.com/victorrub/dns-reset/infra/errors"
)

// networkService is the domain service for Network operations
type networkService struct {
	svc *Service
}

func NewNetworkService(svc *Service) contract.NetworkService {
	return &networkService{
		svc: svc,
	}
}

// ResetLocation creates and activates a new network location
func (s *networkService) ResetLocation(name string) (err error) {

	err = s.svc.net.CreateLocation(name)
	if err != nil {
		return errors.Wrap(err)
	}

	err = s.svc.net.SwitchLocation(name)
	if err != nil {
		return errors.Wrap(err)
	}

	currentLocation, err := s.svc.net.GetCurrentLocation()
	if err != nil {
		return errors.Wrap(err)
	}

	if currentLocation != name {
		message := fmt.Sprintf(domain.ErrCurrentLocation.Error(), currentLocation)
		return errors.NewApplicationError("network.ResetLocation", message)
	}

	return nil
}

// CheckConnection .
func (s *networkService) CheckConnection() (err error) {

	domains := []string{"google.com", "youtube.com", "github.com", "aws.amazon.com"}

	err = s.svc.net.CheckConnection(domains, domain.MaxPingRequests)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}
