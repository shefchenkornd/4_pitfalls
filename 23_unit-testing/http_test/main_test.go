package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HttpTestCase struct {
	name     string      // Название теста
	number   interface{} // Значение, которое будет передаваться в http запрос
	expected []byte      // http ответ
}

// Тестовый сценарий для http запроса
var HttpCases = []HttpTestCase{
	{
		name:     "first test",
		number:   1,
		expected: []byte("1"),
	},
	{
		name:     "second test",
		number:   3,
		expected: []byte("6"),
	},
	{
		name:     "third test",
		number:   5,
		expected: []byte("120"),
	},
	{
		name:     "fourth test",
		number:   0,
		expected: []byte("Send number from 1 to 10"),
	},
	{
		name:     "fifth test",
		number:   11,
		expected: []byte("Send number from 1 to 10"),
	},
	{
		name:     "fifth test",
		number:   "b",
		expected: []byte("Invalid query param."),
	},
}

func TestExecFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandlerExecFactorial)

	for _, test := range HttpCases {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder() // Куда писать Response
			handlerAddr := fmt.Sprintf("http://localhost:8000/factorial?num=%d", test.number)
			request, err := http.NewRequest(http.MethodGet, handlerAddr, nil) // Какой будет *Request
			if err != nil {
				t.Error(err)
			}

			handler.ServeHTTP(recorder, request) // выполняем запрос и ответ записываем в recorder
			// fmt.Printf("Request number: %d, Response factorial: %v\n\n", test.number, recorder.Body.String()) // выводим ответ в консоль

			if recorder.Body.String() != string(test.expected) {
				t.Errorf("Test `%s` failed: input: %v result: %v, expected: %v",
					test.name,
					test.number,
					recorder.Body.String(),
					string(test.expected),
				)
			}
		})
	}
}
