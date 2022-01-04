package main

import "fmt"

type DataService struct {
	Next Service
}

func (device *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Data not update")
		return
	}

	fmt.Println("Data save")
}

func (device *DataService) SetNext(svc Service) {
	device.Next = svc
}



