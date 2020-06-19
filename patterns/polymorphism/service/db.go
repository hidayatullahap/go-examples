package service

import "fmt"

// Implementation of GetData from connection to "database"
type Db struct{}

func (s *Db) GetData() (data Data) {
	// process to get data from database
	fmt.Println("🤖 Beep boop fetch data from database")
	data.Name = "I'm from database"
	data.Age = 2

	return
}

func NewDbService() *Db {
	return &Db{}
}
