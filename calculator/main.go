package main

import (
	"calculator/mathops"
	_ "calculator/registration"
	f "fmt"
)

func main() {

	f.Println(mathops.Add(1, 2))
	f.Println(mathops.Multiply(3, 231))
}
