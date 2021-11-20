package main

import "io"

// io.Reader
// Интерфейс io.Reader предназначен для считывания данных.
// Метод Read возвращает общее количество считанных байт из среза байт и информацию об ошибке, если она возникнет. Если в потоке больше нет данных, то метод должен возвращать ошибку типа io.EOF.
type phoneReader string

func (ph phoneReader) Read(p []byte) (n int, err error) {
	count := 0
	for a, _ := range ph {
		if ph[a] >= 48 && ph[a] <= 57 {
			p[count] = ph[a]
			count++
		}

	}
	return count, io.EOF
}

func main() {
	ph := phoneReader("+1(234)567 9010")
	sl := make([]byte, len(ph))

	ph.Read(sl)

	println(string(sl))
}
