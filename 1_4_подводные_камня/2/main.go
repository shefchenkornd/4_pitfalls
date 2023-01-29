package main

import "fmt"

func main() {
	importantVariableValue := getImportantVariableValue()
	fmt.Println(importantVariableValue)
}


/**
Затененные переменные нередко приводят к ошибкам и возникают при повторном использовании имени переменной в более узкой области.
Когда для объявления переменной в Go задействуется:=, она объявляется внутри области, ограниченной {}.
Но допустим, то же имя переменной используется в более узкой области, например в операторе if.
А причина в повторном объявлении importantVariable в операторе if в этой строке:
*/
func getImportantVariableValue() string {
	importantVariable := "Default Value"
	if true {
		importantVariable := "Non-Default Value"
		fmt.Printf("Important Variable updated to: %s\n", importantVariable)
	}
	return importantVariable
}
