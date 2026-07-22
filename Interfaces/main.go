package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

// Структура Dog
type Dog struct {
	Name string
}

// Dog реализует Speak — автоматически становится Speaker
func (d Dog) Speak() string {
	return "Гав! Я " + d.Name
}

// Структура Cat
type Cat struct {
	Name string
}

// Cat тоже реализует Speak
func (c Cat) Speak() string {
	return "Мяу! Я " + c.Name
}

// Type switch позволяет проверить несколько типов последовательно.
func printType(i any) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %q\n", v)
    case bool:
        fmt.Printf("bool: %t\n", v)
    case Dog:
        fmt.Printf("собака: %s\n", v.Name)
    default:
        fmt.Printf("неизвестный тип: %T\n", v)
    }
}

func main() {
	var s Speaker

	s = Dog{Name: "Бобик"} // Dog подходит
	fmt.Println(s.Speak())

	s = Cat{Name: "Мурка"} // Cat подходит
	fmt.Println(s.Speak())
	fmt.Println("_______________________________________")
	fmt.Println("Извлечение конкретных значений из интерфейса")
	fmt.Println("Объявляем переменную типа any ")
	var i any = 42
	fmt.Printf("i=%v, тип переменно %T\n", i, i)
	num := i.(int)
	fmt.Printf("После чего извлекаем число как int, num=%d, типа %T\n", num, num)
    fmt.Println("_______________________________________")
	// С проверкой (безопасный вариант)
	fmt.Println("Можно присваивать с проверкой")
	value, ok := i.(string)
	if ok {
		fmt.Println("Это строка:", value)
	} else {
		fmt.Println("Это не строка")
	}
	fmt.Println("_______________________________________")
	fmt.Println("Проверка нескольких типов последовательно ")
	var st any = "Строка"
	var bo any = false
	printType(i)
	printType(st)
	printType(bo)
	fmt.Println("_______________________________________")
}
	