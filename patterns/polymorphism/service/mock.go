package service

import "fmt"

// Implementation of GetData from local app
type Mock struct{}

func (s *Mock) GetData() (data Data) {
	// process to get data from local
	fmt.Println("🤖 Beep boop get data from mock")
	data.Name = "I'm from Mock data"
	data.Age = 4

	return
}

func NewMockService() *Mock {
	return &Mock{}
}
