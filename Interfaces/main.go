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

func main() {
	var s Speaker

	s = Dog{Name: "Бобик"} // Dog подходит
	fmt.Println(s.Speak())

	s = Cat{Name: "Мурка"} // Cat подходит
	fmt.Println(s.Speak())
}
