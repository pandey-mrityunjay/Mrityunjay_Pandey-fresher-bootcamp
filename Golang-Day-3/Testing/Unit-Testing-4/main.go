package main

import (
	"fmt"
	service "mocking-example/services"
)

// DataProcessor uses DataService to get and process data
type DataProcesssor struct {
	services service.DataService
}

func (dp *DataProcesssor) ProcessDaata(id int) (string, error) {
	data, err := dp.services.GetData(id)
	if err != nil {
		return "", err
	}
	return "Processed:" + data, nil
}

func main() {
	fmt.Println("Run `go test` to see mocking in action.")
}
