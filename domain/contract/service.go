package contract

// NetworkService holds network operations
type NetworkService interface {
	ResetLocation(name string) (err error)
	CheckConnection() (err error)
}
