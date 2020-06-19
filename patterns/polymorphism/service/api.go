package service

import "fmt"

// Implementation of GetData from connection to "web services"
type Api struct{}

func (s *Api) GetData() (data Data) {
	// process to get data from api
	fmt.Println("🤖 Beep boop fetch data from http server")
	data.Name = "I'm from API"
	data.Age = 1

	return
}

func NewApiService() *Api {
	return &Api{}
}
