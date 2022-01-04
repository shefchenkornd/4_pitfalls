package main

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool	// Флаг, который отвечает за то выполнился ли приём данных
	UpdateSource bool	// Флаг ставит сервис, который обработал данные
}
