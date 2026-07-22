package main

import (
	"fmt"
	"errors"
)


func main() {
	fmt.Println("_______________________________________")
	fmt.Println("Создаём базовую ошибку, со своим сообщением (встроенную)")
	err1 := errors.New("что-то пошло не так")
	fmt.Println(err1) // "что-то пошло не так"
	fmt.Println("_______________________________________")
	fmt.Println("Создаём базовую ошибку, с данными, получеными извне")
	filename := "input.txt"
	err2 := fmt.Errorf("файл %s не найден", filename)
	fmt.Println(err2)
}
	