package main

// В данном примере метод Area определен для структуры Rectangle.
// Если бы мы попытались определить еще один метод с именем Area для этой же структуры с другой сигнатурой, это привело бы к ошибке компиляции.

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5.0, Height: 3.0}
	area := rect.Area()
	fmt.Println("Площадь прямоугольника:", area)
}
