package main

import (
	"fmt"
	"os"
)

var (
	a1 int = 5
	b1 int = 3
)
var a2, b2 int = 10, 20
var a3, b3 int
var af, bf float64 = 11.2, 21.3
var af0, bf0 float64

func sum_1(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на 0")
	}
	return a / b, nil
}

func divideint(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на 0")
	}
	return float64(a) / float64(b), nil
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Определяем функцию add
func add(a int, b int) int {
	return a + b
}

// Передача функций как параметров
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func add2(a, b int) int { return a + b }
func mul(a, b int) int  { return a * b }

// Замыкание - Анонимная ф-ция содержащая внешнию переменную
func makeCounter() func() int {
	count := 0
	return func() int {
		count++ // захватывает count из внешней функции
		return count
	}
}

// отложенный вызов
func readFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // гарантированно закроет файл при выходе из функции

	// работа с файлом...
	return nil
}

// несколько `defer`,  выполняются в порядке **LIFO** (Last In, First Out)
func LIFO_defer() {
	defer fmt.Println("первый defer")
	defer fmt.Println("второй defer")
	defer fmt.Println("третий defer")
	fmt.Println("обычный код")
}
func defer_features() {
	x := 10
	defer fmt.Println("x в defer:", x) // вычислится сейчас, x = 10
	x = 20
	fmt.Println("x в main:", x) // x = 20
}

func main() {
	a4 := 1
	b4 := 2
	fmt.Printf("Функция суммы с простым выводом int значения\n", sum_1(a4, b4))
	result, err := divide(af, bf0)
	if err != nil {
		fmt.Printf("Функция возвращающая ошибку и значение (деление для float64 чисел) выводит ошибку: %v\n", err)
	} else {
		fmt.Printf("Функция возвращающая ошибку и значение (деление для float64 чисел) выводит результат деления: %f\n", result)
	}

	result, err = divide(af, bf)
	if err != nil {
		fmt.Printf("Функция возвращающая ошибку и значение (деление для float64 чисел) выводит ошибку: %v\n", err)
	} else {
		fmt.Printf("Функция возвращающая ошибку и значение (деление для float64 чисел) выводит результат деления: %f\n", result)
	}
	result, err = divideint(a1, b2)
	if err != nil {
		fmt.Printf("Функция возвращающая ошибку и значение (деление для int чисел) выводит ошибку: %v\n", err)
	} else {
		fmt.Printf("Функция возвращающая ошибку и значение (деление для int чисел) выводит результат деления: %f\n", result)
	}
	x, y := split(12)
	fmt.Printf("Голая ф-ция split(12), выводит 2 значения X=%d и Y=%d\n", x, y)
	// Объявляем переменную типа "функция, принимающая два int и возвращающая int"
	var operation func(int, int) int
	// Присваиваем ей функцию add
	operation = add
	// Теперь можем вызывать функцию через переменную
	result2 := operation(5, 3)
	fmt.Printf("Вывод ф-ции присвоенной переменной: %d\n", result2)
	// Передача функций как параметров
	fmt.Printf("Вывод ф-ции которой в качестве параметра пришла другая ф-ция (Сумма): %d\n", applyOperation(5, 3, add2))       // 8
	fmt.Printf("Вывод ф-ции которой в качестве параметра пришла другая ф-ция (Произведение): %d\n", applyOperation(5, 3, mul)) // 15

	// Анонимные функции
	// Присваивание переменной
	increment := func(x int) int {
		return x + 1
	}
	fmt.Printf("Вывод анонимной функции: %d\n", increment(5)) // 6
	// Немедленный вызов
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(10, 20) // сразу вызываем с аргументами
	fmt.Printf("Вывод ф-ции которой сразу с объявлением был передан аргумент (Всё ещё анонимная ф-ция): %d\n", max) // 20

	// Объявляем первую переменную
	counter1 := makeCounter()
	fmt.Printf("Вызов ф-ции счетчика, с внешней переменной с каждым вызовом переменной, ф-ция будет увеличивать её значение на 1, 1шаг: %d\n", counter1())  // 1
	fmt.Printf("Вызов ф-ции счетчика, с внешней переменной с каждым вызовом переменной, ф-ция будет увеличивать её значение на 1, 2 шаг: %d\n", counter1()) // 2

	// Объявляем вторую переменную
	counter2 := makeCounter()
	fmt.Printf("Вызов ф-ции счетчика, с новой внешней переменной, 1шаг: %d\n", counter2())                                                   // 1
	fmt.Printf("Вызов ф-ции счетчика, с старой внешней переменной для демонстрации того, что переменная не забыта, 3 шаг: %d\n", counter1()) // 3

	fmt.Println("Сейчас будет вывод ф-ции LIFO_defer, демонстрирующий LIFO порядок вывода defer")
	LIFO_defer()
	fmt.Println("Вывод закончился")

	fmt.Println("Аргументы отложенной функции вычисляются в момент объявления `defer`, а не в момент выполнения, демонстрация на defer_features")
	defer_features()
	fmt.Println("Как это работает:\n1.Анонимная функция func() { ... } создается\n2.Она замыкает переменную x из внешней области видимости\n3.Функция запоминает ссылку на переменную x, а не её значение на момент создания\n4.Когда функция выполняется (после x = 20), она использует текущее значение x")

}
