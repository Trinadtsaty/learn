package main

import (
    "Methods/bank"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}
type MyInt int

// Метод IsEven — проверяет, чётное ли число
func (mi MyInt) IsEven() bool {
    return mi%2 == 0
}

// Метод Double — возвращает удвоенное значение (тип тот же)
func (mi MyInt) Double() MyInt {
    return mi * 2
}
func (mi *MyInt) DoubleSave() {
    *mi = *mi * 2
}

func (p Person) Greet() string {
    return "Hello, I'm " + p.Name
}
// Создаёт копию объекта
func (p Person) SetAgeWrong(age int) {
    p.Age = age  // меняется только копия
}
// Получает указатель на исходный объект
func (p *Person) SetAge(age int) {
    p.Age = age  // изменяет исходный объект
} 

var pt Person         // значение

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Printf("вывод программы %v\n",p)
    fmt.Println("_______________________________________")
    alice := Person{Name: "Alice", Age: 30}
    fmt.Println("Метод получает копию объекта")
    fmt.Println("Изменили значение возроста на 35")
    alice.SetAgeWrong(35)
    fmt.Printf("Но глобальное значение осталось прежним: %v\n", alice.Age)// всё ещё 30 — не изменилось!
    fmt.Println("_______________________________________")
    fmt.Println("Метод получает ссылку на сам объект")
    fmt.Println("Изменили значение возроста на 35")
    alice.SetAge(35)
    fmt.Printf("Глобальное значение изменилось: %v\n", alice.Age)
    fmt.Println("_______________________________________")
    fmt.Println("Берем переменную типа 'Значение'")
    fmt.Println("Изменили значение возроста на 35, методом SetAge")
    pt.SetAge(35)         // OK: Go делает &p, потому что SetAge на *Person
    fmt.Printf("Глобальное значение изменилось: %v\n", alice.Age)
    fmt.Println("Изменили значение возроста на 40, методом SetAgeWrong")
    alice.SetAgeWrong(40)
    fmt.Printf("Но глобальное значение осталось прежним: %v\n", alice.Age)
    fmt.Println("_______________________________________")
    fmt.Println("Берем переменную типа 'Указатель', она указывает на перемменую alice")
    ptr := &alice
    fmt.Printf("Указатель совпадает со значение: ptr.Age =%v, alice.Age=%v\n", ptr.Age, alice.Age)
    fmt.Println("Меняем значение указателя")
    ptr.SetAge(60)
    fmt.Printf("Переменная, на которую указывал указатель так же поменялась:\nptr.Age =%v, \nalice.Age=%v\n", ptr.Age, alice.Age)
    fmt.Println("_______________________________________")
    fmt.Println("Можно создавать пользовательские типы данных на основе существующих\nЧтобы, применять к ним методы")
    var mi MyInt = 3
    num := MyInt(42)
    fmt.Printf("Был создан тип данных MyInt=%v на основе int\nДля него был определен методы\n", mi)
    
    fmt.Printf("Вызываем метод IsEven для mi=%v, Результат: %v\n", mi, mi.IsEven()) // false
    fmt.Printf("Вызываем метод Double для mi=%v, Результат: %v\n", mi, mi.Double()) // 84
    fmt.Printf("Вызываем метод IsEven для num=%v, Результат: %v\n", num, num.IsEven()) // true
    fmt.Printf("Вызываем метод Double для num=%v, Результат: %v\n", num, num.Double()) // 84
    fmt.Println("_______________________________________")
    fmt.Printf("Вызываем метод DoubleSave для num=%v, ", num) // 84
    num.DoubleSave()
    fmt.Printf("Результат: %v\n", num) // 84   
    fmt.Println("_______________________________________")
    fmt.Println("Импортированный тип данных и методы")
    acc := bank.NewAccount()
    fmt.Printf("Создаём переменную типа NewAccount из импортированого пакета acc=%v,\n", acc)
    acc.Deposit(100.50)
    fmt.Printf("Записываем в переменную значение через метод, acc=%v,\n", acc)
    fmt.Println("Вызываем встроенный метод, который показывает баланс")
    fmt.Printf("Баланс: %.2f\n", acc.GetBalance())
    fmt.Println("_______________________________________")

}