package main

import (
	"fmt"
	"math"
)

// Интерфейсы в языке программирования Go представляют собой абстрактные типы данных,
// которые определяют набор методов, которые должен реализовать любой тип данных, чтобы удовлетворить этот интерфейс.
// Интерфейсы позволяют абстрагировать детали реализации и создавать более гибкий и расширяемый код.

// В Go интерфейс определяется с помощью ключевого слова interface, за которым следует список методов без их реализации.

// Пример определения интерфейса:
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Пример реализации интерфейса Shape для структуры Circle:
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printAreaAndPerimeter(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

// Пример реализации интерфейса Shape для структуры Circle:
func main() {

	circle := Circle{Radius: 5}
	printAreaAndPerimeter(circle)
}
