package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInterface interface{}

	myInterface = 42
	// myInterface = "Hello, World!"
	// myInterface = true
	// myInterface = make(chan int)

	// Определение типа переменной в рантайме
	switch reflect.TypeOf(myInterface).Kind() {
	case reflect.Int:
		fmt.Println("Это int")
	case reflect.String:
		fmt.Println("Это string")
	case reflect.Bool:
		fmt.Println("Это bool")
	case reflect.Chan:
		fmt.Println("Это канал")
	default:
		fmt.Println("Неизвестный тип")
	}
}
