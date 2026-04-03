package main

import (
	"fmt"
)
func addElement(m map[string]int, key string, val int) {
	fmt.Printf("Внутри ф-ции получили map-у m=%v\nПрисвоили ей значение %v по ключу %v\n", m, val, key)
    m[key] = val
	fmt.Printf("Внутри ф-ции map-а после присвоения m=%v\n", m)
}

// 4. Нулевая мапа (nil map)
var m4 map[string]int // nil, len=0, читать можно (возвращает zero value), писать нельзя!

func main() {
	// 1. Через make — пустая мапа (готова к использованию)
	fmt.Println("Структура map-ы (объяявления) make(map[тип данных ключа]тип данных значения)")
	m1 := make(map[string]int)
	fmt.Printf("Пустая map-а m1=%v\n", m1)
	fmt.Println("___")
	// 2. С подсказкой ёмкости (оптимизация)
	m2 := make(map[string]int, 100) // изначальное количество бакетов подберётся под ~100 элементов
	fmt.Printf("Пустая map-а, с кол-вом бакетов необходимых для 100 элементов (1 бакет - 8 элементов), т.е. это 13 бакетов m2=%v\n", m2)
	fmt.Println("___")
	// 3. Литерал
	m3 := map[string]int{
		"alice": 30,
		"bob":   25,
	}
	fmt.Printf("Заполненая map-а, m3=%v\n", m3)
	fmt.Println("___")

	fmt.Printf("Пустая map-а дописывать нельзя, m4=%v\n", m4)
	m1["123"] = 1
	// m4["key"] = 42        // panic: assignment to entry in nil map
	fmt.Println("___")
	fmt.Println("Чтение и запись")
	m := make(map[string]int)
	fmt.Printf("Создали пустую мапу m=%v\n", m)
	// Запись
	m["answer"] = 42
	fmt.Printf("Добавляем запись в map-у m=%v\n", m)
	// Чтение с проверкой существования
	value, ok := m["answer"]
	fmt.Printf("Получаем значение из map-ы, ok=%v, значение=%d\n", ok, value)
	if ok {
		fmt.Println("значение:", value)
	}
	// Чтение без проверки — если ключа нет, вернёт zero value (для int это 0)
	value2, ok2 := m["missing"] // 0
	// Но 0 может быть и реальным значением, поэтому всегда используйте ok, если ноль допустим.
	fmt.Printf("Получаем значение которого не существует из map-ы, ok=%v, значение=%d\n", ok2, value2)
	// Количество элементов
	fmt.Printf("Получаем длинну map-ы len(m)=%d\n",len(m))
	// удаление
	delete(m, "answer")   // удаляет пару; если ключа нет — ничего не делает

	// Проверка после удаления
	if _, ok := m["answer"]; !ok {
		fmt.Println("ключ удалён")
	}
	fmt.Printf("Получаем длинну map-ы после удаления len(m)=%d\n",len(m))
	fmt.Println("___")
	// Итерация по мапе
	m5 := map[string]int{"a": 1, "b": 2, "c": 3}

	for key, value := range m5 {
		fmt.Printf("Перебираем в массиве значения ключ=%s, значение=%d\n",key, value)
	}

	// Только ключи
	for key := range m5 {
		fmt.Printf("Берем только ключи=%s \n",key)
	}

	// Важно: порядок итерации случаен (рандомизирован). Не полагайтесь на порядок!
	// Причина: защита от хеш-дозирования (hash flooding) — злоумышленник не может предсказать порядок.
	fmt.Println("___")
	fmt.Println("map-ы как сылочный тип")
	// map-ы это ссылочный тип, а значит если изменить map-у внутри ф-ции, изменения сохранятся и внешне
	m6 := make(map[string]int)
	fmt.Printf("Создали пустую map-у m6=%v\n",m6)
    addElement(m6, "foo", 42)
	fmt.Printf("map-а после ф-ции m6=%v\n",m6)
    // fmt.Println(m6["foo"]) // 42
	fmt.Println("___")
}
