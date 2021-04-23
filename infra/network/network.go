package network

// Interface set methods to interact with Mac OS Network
type Interface interface {
	// networksetup operations
	GetCurrentLocation() (location string, err error)
	ListLocations() (locations []string, err error)
	SwitchLocation(name string) (err error)
	CreateLocation(name string) (err error)
	DeleteLocation(name string) (err error)

	// ping operations
	CheckConnection(domains []string, requestLimit int) (err error)
}

// Communicator returns the methods to interact with Mac OS Network operations
type Communicator struct{}

// NewCommunicator return the instance of Communicator
func NewCommunicator() Interface {
	return &Communicator{}
}
