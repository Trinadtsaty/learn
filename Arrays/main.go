package main

import (
	"fmt"
)

var arr [5]int
var matrix [3][3]int

// func getElement(slice []int, index int) (int, error) {
// 	if index < 0 || index >= len(slice) {
// 		return -1, errors.New("индекс вне границ слайса")
// 	}
// 	return slice[index], nil
// }

func modifyArray(arr [3]int) {
    arr[0] = 100
    fmt.Println("Внутри функции:", arr) // [100, 20, 30]
}

func modifyArrayPtr(arr *[3]int) {
    arr[0] = 100  // автоматическое разыменование (*arr)[0] = 100
    fmt.Println("Внутри функции:", *arr)
}


func main() {
	fmt.Println("Hello World")
	fmt.Println(arr)

	arr1 := [3]int{1, 2, 3}           // массив из 3 элементов
	arr2 := [5]int{1, 2, 3}           // первые три элемента: 1,2,3, остальные: 0,0
	arr3 := [...]int{1, 2, 3}         // ... — компилятор сам подсчитает длину (3)
	arr4 := [5]int{0: 10, 2: 30}      // индексы 0 и 2 инициализированы, остальные — 0
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	arr := [3]int{10, 20, 30}
	arr[0] = 100          // изменение элемента
	value := arr[1]       // чтение: 20

	fmt.Println(len(arr)) // 3 — длина массива (константа времени компиляции)
	fmt.Printf("Вывод массива %v\n", arr)
	fmt.Printf("Вывод первого элемента, присвоенного переменной %d", value)


	original1 := [3]int{10, 20, 30}
    modifyArray(original1)
    fmt.Println("После вызова:", original1) // [10, 20, 30] — не изменился!

	original2 := [3]int{10, 20, 30}
    modifyArrayPtr(&original2)
    fmt.Println("После вызова:", original2) // [100, 20, 30] — изменился!


	// arr := [3]int{1, 2, 3}
	// arr[5] = 10  // panic: index out of range [5] with length 3
	// answer, err :=getElement(arr, 6)
	// if err != nil {
	// 	fmt.Printf("ошибок нет, вот индекс %d\n", )
	// } else {
	// 	fmt.Printf("У нас возникла ошибка %s\n", )
	// }
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	fmt.Println(a == b) // true
	fmt.Println(a == c) // false

	fmt.Printf("Матрица 3*3 %v\n", matrix)
	matrix1 := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
	}
	fmt.Printf("Матрица заданная ручками 2*3 %v\n", matrix1)
	// Доступ
	matrix[0][1] = 10  // изменяет элемент во второй строке, втором столбце
	// Длина
	fmt.Printf("Количество строк: %d\n", len(matrix))
	fmt.Printf("Количество столбцов: %d\n", len(matrix[0]))

	
}


