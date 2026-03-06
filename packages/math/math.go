package math

import f "fmt"

// Add — экспортируемая функция (доступна извне)
func Add(a, b int) int {
    logOperation("Add") // вызываем внутреннюю функцию
    return a + b
}

// sub — неэкспортируемая функция (только внутри пакета)
func sub(a, b int) int {
    logOperation("sub")
    return a - b
}

// logOperation — неэкспортируемая вспомогательная функция
func logOperation(op string) {
    f.Printf("Выполняется операция: %s\n", op)
}