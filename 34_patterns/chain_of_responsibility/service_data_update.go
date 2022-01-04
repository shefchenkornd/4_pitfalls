package main

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (device *UpdateDataService) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data in service [%s] already update.\n", device.Name)
		device.Next.Execute(data)
		return
	}

	fmt.Printf("Update data from service [%s].\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *UpdateDataService) SetNext(svc Service) {
	device.Next = svc
}



