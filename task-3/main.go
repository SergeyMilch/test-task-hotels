package main

import (
	"fmt"
	"math"
)

// isSimpleNumber проверяет, является ли число простым
func isSimpleNum(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true // 2 - это единственное четное простое число
	}
	if n%2 == 0 {
		return false // Исключаем все четные числа
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 { // не проверяем делители, большие квадратного корня
		if n%i == 0 {
			return false
		}
	}
	return true
}

// simpleNumber возвращает массив простых чисел в заданном диапазоне
func simpleNumber(min int, max int) []int {
	if min > max {
		return nil // Если минимальное значение больше максимального, возвращаем nil
	}
	var simpleNums []int
	for number := min; number <= max; number++ {
		if isSimpleNum(number) {
			simpleNums = append(simpleNums, number)
		}
	}
	return simpleNums
}

func main() {
	rangeNumbers := simpleNumber(11, 20)

	fmt.Println(rangeNumbers)
}
