package main

import (
	"fmt"
	"sort"
)

func main() {
	// Исходная последовательность температурных колебаний
	temperatures := []float64{-25.0, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Сортируем температуры по возрастанию
	sort.Float64s(temperatures)

	// Создаем карту (map) для хранения групп с шагом в 10 градусов
	groups := make(map[int][]float64)

	// Итерируемся по отсортированной последовательности и добавляем значения в соответствующие группы
	for _, temperature := range temperatures {
		groupKey := int(temperature/10) * 10 // Определяем ключ группы для текущей температуры
		groups[groupKey] = append(groups[groupKey], temperature)
	}

	// Создаем срез для хранения ключей групп и сортируем его
	keys := make([]int, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Выводим группы в порядке возрастания ключей
	for _, key := range keys {
		fmt.Printf("%d: {", key)
		group := groups[key]
		for i, temperature := range group {
			fmt.Printf("%.1f", temperature)
			if i < len(group)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("}\n")
	}
}
