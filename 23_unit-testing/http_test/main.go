package main

import (
	"log"
	"net/http"
	"strconv"
)

const (
	MinValue = 1
	MaxValue = 10
)

// Совет: необходимо кэшировать повторные запросы, если у вас есть квадратичная сложность — O(n2) в вычислениях!
// Квадратичная временная сложность — это когда производительность алгоритма прямо пропорциональна размеру входных данных в квадрате.
// Проще говоря, это линейная временная сложность в квадрате.
// Если, к примеру, в нашем наборе данных есть 2 элемента, за время работы алгоритма будет выполнено 4 операции.
// Если в наборе 4 элемента, то операций будет 16. При 6 элементах будет 36 операций, и так далее.

var cache map[int]int // Сюда будем кэшировать наши запросы по факториалу

func init()  {
	cache = make(map[int]int)
}

// Создаём простой http server для вычисления factorial и пишем на него http-tests
func main() {
	http.HandleFunc("/factorial", HandlerExecFactorial)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

// HandlerExecFactorial хендлер по обработке http запроса
func HandlerExecFactorial(w http.ResponseWriter, req *http.Request) {
	param := req.URL.Query().Get("num")
	// -or-
	// param := req.FormValue("num")
	num, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid query param."))
		// -or-
		// http.Error(w, "Invalid query param.", http.StatusBadRequest)
		return
	}

	if MinValue > num || num > MaxValue {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Send number from 1 to 10"))
		// -or-
		// http.Error(w, "Send number from 1 to 10", http.StatusBadRequest)
		return
	}

	result := factorial(num)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(result)))
}

// factorial функция, определённая на множестве неотрицательных целых чисел.
// Формула: n! = 1 * 2 * ... * n
// Например: 5 ! = 1 * 2 * 3 * 4 * 5 = 120.
func factorial(n int) int {
	var result = 1

	if val, ok := cache[n]; ok {
		return val
	}
	
	for i := 1; i <= n; i++ {
		result *= i
	}
	// сохраняем в кэш
	cache[n] = result

	return result
}
