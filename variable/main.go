package main

import (
    "fmt"
	"unsafe"
)
// '''
// dsdsdsd
// '''
var age int = 30
var price float64 = 19.99
var active bool = true
var message string = "Hello"
var count int       // 0
var isReady bool    // false
var text string     // ""
var x, y int = 10, 20
var (
    name   string = "Alice"
    height    int    = 25
    weight float64
)
var i int = 42
var f float64 = float64(i)   // int → float64
var u uint = uint(f)         // float64 → uint (дробная часть отбрасывается)

var original string = "hello"
var bytes []byte = []byte(original)  // string → []byte (копирует данные)
var back string = string(bytes)      // []byte → string (тоже копирует)

// Константы
const pi = 3.14159
const appName string = "MyApp"

func main() {
    var i int
    var ui uint
	user_name := "Bob"      // string
    user_age := 42          // int
    user_price := 99.5      // float64
    user_enabled := true    // bool
	fmt.Printf("Привет %s, я знаю что тебе %d лет и ты работаешь за %f.\nТвой статус завявки: %v",user_name, user_age, user_price, user_enabled)	
    // result := fmt.Sprintf("Привет %s, я знаю что тебе %d лет и ты работаешь за %f.\nТвой статус заявки: %v", user_name, user_age, user_price, user_enabled)
	// fmt.Println(result)
	fmt.Println()
    fmt.Printf("Размер int: %d байт\n", unsafe.Sizeof(i))   // 8 на 64-битной системе
    fmt.Printf("Размер uint: %d байт\n", unsafe.Sizeof(ui)) // тоже 8
	fmt.Println()
	fmt.Printf("Константы числовые: %v, %f;\nТекстовая: %s", pi, pi, appName)
	fmt.Println()
    const (
        Sunday = (iota)  // 0
        Monday         // 1
        Tuesday        // 2
        Wednesday      // 3
        Thursday       // 4
        Friday         // 5
        Saturday       // 6
    )
    const (
        _ = iota           // пропускаем 0
        KB = 1 << (10 * iota)  // 1 << 10 = 1024
        MB                     // 1 << 20 = 1048576
        GB                     // 1 << 30 = 1073741824
    )
    fmt.Printf(
        "Sunday=%d;\nMonday=%d;\nTuesday=%d;\nWednesday=%d;\nThursday=%d;\nFriday=%d;\nSaturday=%d\n",
        Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
    fmt.Println()
    fmt.Printf("KB=%d, MB=%d, GB=%d\n", KB, MB, GB)
    fmt.Println()

}

