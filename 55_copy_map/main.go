package main

import "fmt"

type counter struct {
	Properties map[string]string
}

// Скопировать map просто так нельзя!
// Нужно сделать шаги, показанные ниже:
func main() {
	c := counter{
		Properties: map[string]string{
			"a": "A",
			"b": "B",
			"c": "C",
		},
	}

	fmt.Println("Original map:", c.Properties)

	copiedMap := make(map[string]string)
	copiedMap = copyMap(c.Properties, copiedMap)

	fmt.Println("Copied map:", copiedMap)

	// Попытаемся изменить исходную map, чтобы проверить, что скопированная map осталась неизменной:
	c.Properties["d"] = "D"
	fmt.Println("Original map:", c.Properties)
	fmt.Println("Copied map:", copiedMap)
}

func copyMap(src, dst map[string]string) map[string]string {
	for k, v := range src {
		dst[k] = v
	}

	return dst
}
