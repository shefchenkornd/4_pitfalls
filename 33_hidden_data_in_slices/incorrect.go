package main

import "fmt"

// При перенарезке получившийся слайс будет ссылаться на массив исходного слайса. Не забывайте об этом.
// Иначе может возникнуть непредсказуемое потребление памяти, когда приложение разместит в ней крупные временные слайсы
// и создаст из них новые, чтобы ссылаться на небольшие куски исходных данных.
// Чтобы избежать этой ошибки, удостоверьтесь, что копируете нужные данные из временного слайса (вместо перенарезки).
func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 0xc000102000

	result := raw[:3]
	return result
}

func main() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0]) // 3 10000 0xc000102000
}
