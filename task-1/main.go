package main

import (
	"fmt"
)

// DeclineComputer склоняет слово "компьютер" в зависимости от числа
func declineComputer(number int) string {
	lastDigits := number % 100  // получаем последние две цифры числа
	// исключения для чисел от 11 до 19
	if lastDigits > 10 && lastDigits < 20 {
		return fmt.Sprintf("%d компьютеров", number)
	}

	switch number % 10 {
	case 1:
		return fmt.Sprintf("%d компьютер", number)
	case 2, 3, 4:
		return fmt.Sprintf("%d компьютера", number)
	default:
		return fmt.Sprintf("%d компьютеров", number)
	}
}

func main() {
	fmt.Println(declineComputer(99))
	fmt.Println(declineComputer(25))
	fmt.Println(declineComputer(41))
	fmt.Println(declineComputer(11))
	fmt.Println(declineComputer(19001))
	fmt.Println(declineComputer(1))
	fmt.Println(declineComputer(42))
	fmt.Println(declineComputer(1048))
	fmt.Println(declineComputer(0))
}