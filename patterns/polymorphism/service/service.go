package service

type Data struct {
	Name string
	Age  int
}

// Every service that implement this interface must have GetData
type ContractService interface {
	GetData() Data
}
