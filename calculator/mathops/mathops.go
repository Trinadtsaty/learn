package mathops

import f "fmt"

func init(){
	f.Println("Пакет mathops инициализирован")
}

func Add (a,b int) int {
	logOperation("Add")
	return a + b
}
func Multiply (a,b int) int {
	logOperation("Multiply")
	return a * b
}
func logOperation (op string) {
	f.Println("Выполняется операция: ", op)
}


// `Add(a, b int) int` и `Multiply(a, b int) int`