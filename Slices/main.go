package main

import (
	"fmt"
)

var nil_s []int 
var sl2 [5]int
var sl3 []int  

func main() {
	fmt.Printf("%s\n","Hello World")

	s := []int{1, 2, 3}  // len=3, cap=3

	s1 := make([]int, 5)      // len=5, cap=5, элементы инициализированы нулями
	s2 := make([]int, 3, 5)   // len=3, cap=5, доступно 3 элемента, резерв на 2

	arr := [5]int{10, 20, 30, 40, 50}
	s3 := arr[1:4]  // [20, 30, 40], len=3, cap=4 (от индекса 1 до конца массива)

	fmt.Printf("Общий слайс %v, его len=%d, его cap=%d\n",s,len(s),cap(s))
	fmt.Printf("Элементы инициализированы нулями %v, его len=%d, его cap=%d\n",s1,len(s1),cap(s1))
	fmt.Printf("Доступно 3 элемента, резерв на 2 %v, его len=%d, его cap=%d\n",s2,len(s2),cap(s2))
	fmt.Printf("Общий слайс %v, его len=%d, его cap=%d\n",arr,len(arr),cap(arr))
	fmt.Printf("Общий слайс %v, его len=%d, его cap=%d\n",s3,len(s3),cap(s3))

	empty_s := []int{}  
	empty_s2 := make([]int, 0)
	fmt.Println("___")
	fmt.Printf("nil slice: %v, его len=%d и его cap=%d;\nНулевой slice: %v, его len=%d и его cap=%d;\nЕщё один нулевой slice: %v, его len=%d и его cap=%d.\n",nil_s, len(nil_s), cap(nil_s), empty_s, len(empty_s), cap(empty_s), empty_s2, len(empty_s2), cap(empty_s2))
	fmt.Println("___")

	arr1 := [5]int{10, 20, 30, 40, 50}
	s4 := arr1[1:4]   // элементы 1, 2, 3 (индексы 1,2,3)
	fmt.Println(s4)  // [20, 30, 40]
	fmt.Println(len(s4), cap(s4)) // 3, 4
	fmt.Println("___")
	s5 := []int{1, 2, 3}
	fmt.Printf("Начальный слайс %v, len=%d, cap=%d;\n", s5, len(s5), cap(s5))
	s5 = append(s5, 4)           // добавляем один элемент
	fmt.Printf("Добавили новый элемент %v, len=%d, cap=%d;\n", s5, len(s5), cap(s5))
	s5 = append(s5, 5, 6, 7)     // добавляем несколько
	fmt.Printf("Добавляем несколько новых элементов %v, len=%d, cap=%d;\n", s5, len(s5), cap(s5))
	s5 = append(s5, []int{8,9}...) // добавляем другой слайс
	fmt.Printf("Добавляем другой слайс %v, len=%d, cap=%d;\n", s5, len(s5), cap(s5))
	fmt.Println("___")
	fmt.Println("Функция copy")
	dst := make([]int, 5)
	src := []int{1, 2, 3}
	n := copy(dst, src)  // n = 3, dst[0:3] = [1,2,3], остальные 0
	fmt.Printf("Добавляем другой что-то %v\n", n)
	fmt.Printf("Добавляем другой что-то %v\n", dst)
	fmt.Printf("Добавляем другой что-то %v\n", src)
	fmt.Println("___")
	sl1 := [5]int{1,2,3}
	fmt.Printf("Разница между слайсом и массивом? %v, len=%d, cap=%d;\n", sl1, len(sl1), cap(sl1))
	fmt.Printf("Разница между слайсом и массивом? %v, len=%d, cap=%d;\n", sl2, len(sl2), cap(sl2))
	fmt.Printf("Разница между слайсом и массивом? %v, len=%d, cap=%d;\n", sl3, len(sl3), cap(sl3))
	sl3 = append(sl3, 1)
	fmt.Printf("Разница между слайсом и массивом? %v, len=%d, cap=%d;\n", sl3, len(sl3), cap(sl3))

}