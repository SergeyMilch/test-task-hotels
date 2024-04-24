package main

import (
	"fmt"
)

// FindCommonDivisors находит общие делители для списка чисел
func findCommonDivisors(nums []int) []int {
	if len(nums) == 0 {
		return nil  // Если массив пустой, возвращаем nil
	}

	// Находим минимальное число в массиве
	minNum := nums[0]

	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}

	var commonDivisors []int

	// Перебираем все числа от 1 до minNum включительно
	for div := 2; div <= minNum; div++ {
		isDivisor := true
		// Проверяем, делит ли div все числа в массиве
		for _, num := range nums {
			if num % div != 0 {
				isDivisor = false
				break  // Если div не делит хотя бы одно число, прерываем внутренний цикл
			}
		}
		// Если div является делителем для всех чисел, добавляем его в срез
		if isDivisor {
			commonDivisors = append(commonDivisors, div)
		}
	}

	return commonDivisors
}

func main() {
	numbers := []int{42, 12, 18}

	fmt.Println("Общие делители чисел", numbers, "-", findCommonDivisors(numbers))
}
