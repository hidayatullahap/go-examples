package main

import (
	"fmt"
	"time"
)

// Get timezone string offset
func main() {
	t := time.Now()

	str := t.Format("-0700")
	fmt.Println(str)

	fmt.Println("Location : ", t.Location(), " Time : ", t)

	location, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
	}

	str2 := t.In(location).Format("-0700")
	fmt.Println(str2)
}

