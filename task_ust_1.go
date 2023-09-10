package main

import (
	"fmt"
	"strings"
	"time"
)

//	Меньшее количество аллокаций памяти: При использовании оператора + для конкатенации строк каждый раз создается новая строка,
//	что приводит к аллокации дополнительной памяти и копированию содержимого предыдущих строк.
//	Это может стать проблемой при объединении большого количества строк или при выполнении операции конкатенации в цикле.
// 	Функция Join, напротив, создает буфер фиксированного размера и копирует содержимое всех строк в этот буфер, избегая лишних аллокаций.
//
//  Оптимизация компилятора: Некоторые компиляторы Go (например, компилятор от Google)
//  могут оптимизировать использование strings.Join для более эффективной работы с памятью и производительностью.
//
//  Код читается легче: Использование strings.Join делает код более ясным и читаемым, потому что намного легче понять
//  намерение конкатенации строк с использованием этой функции, чем с использованием оператора +.

func concatWithPlus(strSlice []string) string {
	result := ""
	for _, s := range strSlice {
		result += s
	}
	return result
}

func concatWithJoin(strSlice []string) string {
	return strings.Join(strSlice, "")
}

func generateLongString() string {
	strSlice := make([]string, 1000) // Создаем срез строк из миллиона элементов
	for i := range strSlice {
		strSlice[i] = "Это"
	}
	return concatWithPlus(strSlice)
}

func measureTimeDifference() {
	longString := generateLongString()

	// Запускаем таймер для функции с оператором +
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		resultPlus := concatWithPlus([]string{longString, longString})
		_ = resultPlus
	}
	elapsedTimePlus := time.Since(startTime)

	// Запускаем таймер для функции с использованием strings.Join
	startTime = time.Now()
	for i := 0; i < 1000; i++ {
		resultJoin := concatWithJoin([]string{longString, longString})
		_ = resultJoin
	}
	elapsedTimeJoin := time.Since(startTime)

	// Выводим результаты и разницу во времени
	fmt.Printf("Время выполнения с использованием оператора +: %s\n", elapsedTimePlus)
	fmt.Printf("Время выполнения с использованием strings.Join: %s\n", elapsedTimeJoin)
	fmt.Printf("Разница во времени выполнения: %s\n", elapsedTimeJoin-elapsedTimePlus)
}

func main() {
	measureTimeDifference()
}
