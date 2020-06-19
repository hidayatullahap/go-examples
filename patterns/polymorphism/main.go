package main

import (
	"fmt"
	"polymorphism/service"
)

// check service packages to see implementation
func main() {
	apiService := service.NewApiService()
	dbService := service.NewDbService()
	mockService := service.NewMockService()

	fmt.Println("running services")

	dataApi := apiService.GetData()
	fmt.Println("fetch success, data: ", dataApi)

	dataDb := dbService.GetData()
	fmt.Println("fetch success, data: ", dataDb)

	dataMock := mockService.GetData()
	fmt.Println("fetch success, data: ", dataMock)
}
