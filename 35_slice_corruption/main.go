package main

import (
	"bytes"
	"fmt"
)

// Задача:
// Допустим, вам нужно переписать путь (хранящийся в слайсе). Чтобы ссылаться на каждую папку, вы его перенарезаете,
// изменяя имя первой папки, а затем комбинируете имена в новый путь.
//
// Решение:
// Так не сработает. Вместо AAAAsuffix/BBBBBBBBB вы получите AAAAsuffix/uffixBBBB.
// Причина в том, что слайсы обеих папок ссылаются на один и тот же массив данных из исходного слайса пути.
// То есть исходный путь тоже изменился. Это может быть проблемой для вашего приложения.
// Ее можно решить, разместив в памяти новые слайсы и скопировав туда нужные данные. Другой выход: использовать полное выражение слайса (full slice expression).
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	// dir1 := path[:sepIndex] // incorrect
	dir1 := path[:sepIndex:sepIndex] // correct | Объяснение: второе двоеточие вводит значение емкости, которое должно быть меньше или равно емкости исходного фрагмента или массива с поправкой на исходную точку.
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) // выводит: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) // выводит: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})

	fmt.Println("path =>", string(path))
	// when INCORRECT, output: AAAAsuffix/uffixBBBB
	// when CORRECT, output: AAAAsuffix/BBBBBBBBB
}
