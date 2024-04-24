package main

import (
	"fmt"
)

// multiplicationTable печатает таблицу умножения до заданного числа
func multiplicationTable(number int) {
	// Печатаем шапку таблицы
	for i := 1; i <= number; i++ {
		fmt.Printf("\t%d", i)
	}
	fmt.Println()

	// Печатаем строки таблицы умножения
	for i := 1; i <= number; i++ {
		fmt.Printf("%d", i)
		for j := 1; j <= number; j++ {
			fmt.Printf("\t%d", i*j)
		}
		fmt.Println()
	}
}

func main() {
	number := 9
	
	multiplicationTable(number)
}
