package main

import "fmt"

// Определяем структуру Human
type Human struct {
	Name  string
	Age   int
	Email string
}

// Метод SayHello для структуры Human
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Определяем структуру Action с встраиванием Human
type Action struct {
	Human
	ActionName string
}

// Метод DoAction для структуры Action
func (a *Action) DoAction() {
	fmt.Printf("%s выполняет действие: %s\n", a.Name, a.ActionName)
}

func main() {
	// Создаем экземпляр структуры Action
	action := Action{
		Human:      Human{Name: "Илья", Age: 20, Email: "ilya@example.com"},
		ActionName: "плавает",
	}

	// Вызываем методы как из Human, так и из Action
	action.SayHello()
	action.DoAction()

	return
}
