package main

import "fmt"

// Функция для нахождения пересечения двух неупорядоченных множеств
func Intersection(set1, set2 []int) []int {
	// Создаем хэш-таблицу для элементов из первого множества
	elementMap := make(map[int]bool)

	// Заполняем хэш-таблицу элементами из первого множества
	for _, elem := range set1 {
		elementMap[elem] = true
	}

	// Создаем список для хранения общих элементов
	intersection := []int{}

	// Проверяем элементы из второго множества на наличие в хэш-таблице
	for _, elem := range set2 {
		if elementMap[elem] {
			intersection = append(intersection, elem)
		}
	}

	return intersection
}

func main() {
	// Пример использования
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := Intersection(set1, set2)

	fmt.Println("Пересечение множеств:", result)
}
