# Начало
**Цель проекта**

Создать собственную минимальную графическую библиотеку для Windows на Go, работающую напрямую через WinAPI (без фреймворков). Библиотека должна позволять создавать окна, обрабатывать события, рисовать примитивы и тексты, а также взаимодействовать с пользователем (кнопки, поля ввода). В перспективе этот код станет основой для пет-проектов, включая десктопное приложение для генерации картинок через OpenRouter.

**Зачем это**

— Понять внутреннее устройство GUI (не использовать чёрные ящики).

— Получить контроль над каждым пикселем и сообщением.

— Научиться работать с системными вызовами, указателями, памятью.

— Создать переиспользуемый модуль для своих будущих проектов.

# Шаг 1
Создание пустого исполняемый файл на Go, подготовлена база для вызовов Windows API.

**Создай файл main.go с таким содержимым:**
```go
package main

import (
    "golang.org/x/sys/windows" // пакет для работы с Windows API
)

func main() {
    
}
```
**Что здесь:**

**package main** — говорит, что это исполняемая программа.

**import "golang.org/x/sys/windows"** — подключаем пакет. Он нужен, потому что в нём есть константы (например, windows.SW_SHOW), структуры (windows.WNDCLASSEXW) и функции (windows.ShowWindow). Без него пришлось бы объявлять всё вручную.

# Шаг 2

Загрузка системной библиотеки user32.dll и получение указателей на функции оконного API. На этом этапе программа получает доступ к функциям Windows, необходимым для создания и управления окнами.

Нужно, чтобы программа могла создать окно и не зависала при работе с ним. 

**Для этого нам нужно:**

- Установить пакет `runtime`. Это стандартный пакет Go, который управляет выполнением программы: горутинами, памятью, потоками ОС. Он всегда доступен, его не надо скачивать. 

```go
import (
    "runtime"
    "golang.org/x/sys/windows"
)
```

- Добавляем `LockOSThread()` Это функция из пакета runtime. Она говорит Go: «привяжи текущую горутину к текущему потоку операционной системы и не отпускай».

```go
func main() {
	runtime.LockOSThread()
}
```

**Почему это обязательно для окон**

Windows устроена так, что все **операции с окнами** (создание, рисование, обработка кликов) должны выполняться **в одном и том же потоке**. Это историческое ограничение WinAPI.

Go по умолчанию может переносить горутину с одного системного потока на другой. Если это случится в тот момент, когда мы создаём окно или обрабатываем сообщения, Windows может решить, что мы обращаемся к окну из неправильного потока, и программа упадёт или зависнет.

Без `LockOSThread()` наше окно может:

- не отображаться,

- не реагировать на клики,

- сразу закрыться с ошибкой.

***Горутина** — легковесный поток в Go, который выполняется поверх системных потоков. Запускается ключевым словом go. В отличие от потоков ОС, горутины имеют меньший накладной расход и управляются рантаймом Go.*

# Шаг 3

Регистрация класса окна. Перед созданием самого окна необходимо зарегистрировать его класс — структуру, описывающую внешний вид и поведение будущего окна, включая ссылку на функцию-обработчик событий.

На этом этапе мы сообщаем Windows о том, что собираемся создавать окна с определёнными свойствами. Для этого необходимо:

1. Объявить константы, которые нужны для обработки сообщений.

2. Загрузить функции из user32.dll, которые нам понадобятся (мы уже загрузили их ранее, но сейчас добавим недостающие).

3. Написать функцию-обработчик сообщений (WndProc).

4. Написать функцию регистрации класса (registerWindowClass).

Начнём по порядку.


## Часть 1. Добавляем недостающие глобальные объявления

В предыдущих шагах мы объявили `user32` и переменные для функций, но не объявили константы `WM_DESTROY` и не получили указатели на `DefWindowProcW` и `PostQuitMessage`. Сейчас мы это сделаем.

Открой свой `main.go` и добавь после импортов следующие строки (они уже были в полном коде, но сейчас мы их вводим поэтапно):

**Объявление:**

```go
// Константы для сообщений Windows
const (
    WM_DESTROY = 0x0002
    WM_QUIT    = 0x0012
)

// Загружаем библиотеку user32.dll
var user32 = windows.NewLazySystemDLL("user32.dll")

// Получаем указатели на функции
var (
    procRegisterClassExW = user32.NewProc("RegisterClassExW")
    procDefWindowProcW   = user32.NewProc("DefWindowProcW")
    procPostQuitMessage  = user32.NewProc("PostQuitMessage")
    // остальные функции добавим позже, когда они понадобятся
)
```

**Пояснение:**

- `WM_DESTROY` и `WM_QUIT` — числовые коды, которые Windows использует для обозначения событий.

- `user32.NewLazySystemDLL` — загружает системную библиотеку `user32.dll`.

- `NewProc` — получает из библиотеки адрес конкретной функции.

Импорты для этого шага: `"unsafe"`, `"syscall"`, `"log"`. Добавь их, если ещё нет.

## Часть 2. Пишем функцию-обработчик WndProc

Создадим пустую функцию с нужной сигнатурой:

```go
//export WndProc
func WndProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
    return 0
}
```

Теперь добавим логику: если пришло сообщение `WM_DESTROY` — вызываем `PostQuitMessage(0)` и возвращаем 0. Иначе передаём управление стандартному обработчику `DefWindowProcW`.

```go
//export WndProc
func WndProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
    switch msg {
    case WM_DESTROY:
        procPostQuitMessage.Call(0)
        return 0
    default:
        ret, _, _ := procDefWindowProcW.Call(hWnd, uintptr(msg), wParam, lParam)
        return ret
    }
}
```

Объяснение работы switch:

- `switch msg` сравнивает значение переменной `msg` с константами.

- Если `msg` равна `2` (WM_DESTROY), то выполняется блок `case`. Windows шлёт это сообщение, когда окно закрывается (нажатие крестика).

- В `case` мы вызываем `PostQuitMessage(0)` — эта функция посылает в очередь сообщений `WM_QUIT`, что приведёт к завершению цикла обработки.

- Если `msg` не равна `2`, то срабатывает `default`, где мы вызываем `DefWindowProcW` — стандартный обработчик Windows, который отвечает за поведение окна по умолчанию (перемещение, изменение размера и т.п.). Возвращаем его результат.

## Часть 3. Функция регистрации класса `registerWindowClass`

Объявим функцию, которая будет регистрировать класс окна. Она возвращает идентификатор класса (атом).

```go
func registerWindowClass() uint16 {
    return 0
}
```

Теперь наполним её шаг за шагом.

### 3.1. Создание имени класса в формате UTF-16

Windows использует двухбайтовую кодировку UTF-16 для строк. Преобразуем строку `"MyGoWindowClass"` в указатель на UTF-16:

```go
className, err := windows.UTF16PtrFromString("MyGoWindowClass")
if err != nil {
    log.Fatal("UTF16PtrFromString failed:", err)
}
```

- `UTF16PtrFromString` возвращает `*uint16` и ошибку. Если ошибка не `nil`, мы завершаем программу.

### 3.2. Заполнение структуры `WNDCLASSEXW`

Создаём структуру, описывающую класс. Все поля обязательны.

```go
wc := windows.WNDCLASSEXW{
    Size:        uint32(unsafe.Sizeof(windows.WNDCLASSEX{})),
    Style:       0,
    WndProc:     syscall.NewCallback(WndProc),
    ClsExtra:    0,
    WndExtra:    0,
    Instance:    0,
    Icon:        0,
    Cursor:      0,
    Background:  0,
    MenuName:    nil,
    ClassName:   className,
    IconSm:      0,
}
```

Пояснение полей (они уже были описаны в предыдущем ответе, но повторю кратко):

- **Size** — размер структуры, Windows проверяет его.

- **Style** — стиль класса (0 = стандартный).

- **WndProc** — указатель на обработчик, получаем через syscall.NewCallback.

- **ClsExtra / WndExtra** — зарезервированы, ставим 0.

- *Instance* — дескриптор экземпляра (0 = текущий).

- **Icon, Cursor, Background** — пока стандартные (0).

- **MenuName** — меню не используем, `nil`.

- **ClassName** — указатель на имя класса.

- **IconSm** — маленькая иконка (0).

### 3.3. Вызов `RegisterClassExW`

Передаём адрес структуры в функцию Windows:

```go
ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
```

- `&wc` — адрес структуры.

- `unsafe.Pointer` — преобразуем в сырой указатель.

- `uintptr` — приводим к целому для передачи.

- Если `ret == 0`, регистрация не удалась — выходим с ошибкой.

```go
if ret == 0 {
    log.Fatal("RegisterClassExW failed")
}
```

### 3.4. Возврат атома класса

Функция возвращает uint16(ret) — это идентификатор зарегистрированного класса.

```go
return uint16(ret)
```

**Итоговый код `registerWindowClass`:**

```go
func registerWindowClass() uint16 {
    className, err := windows.UTF16PtrFromString("MyGoWindowClass")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    wc := windows.WNDCLASSEXW{
        Size:        uint32(unsafe.Sizeof(windows.WNDCLASSEX{})),
        Style:       0,
        WndProc:     syscall.NewCallback(WndProc),
        ClsExtra:    0,
        WndExtra:    0,
        Instance:    0,
        Icon:        0,
        Cursor:      0,
        Background:  0,
        MenuName:    nil,
        ClassName:   className,
        IconSm:      0,
    }

    ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
    if ret == 0 {
        log.Fatal("RegisterClassExW failed")
    }
    return uint16(ret)
}
```

## Часть 4. Проверка в `main`

В функции main() мы уже вызвали runtime.LockOSThread(). Теперь добавим вызов регистрации класса:

```go
func main() {
    runtime.LockOSThread()
    registerWindowClass()
}
```

На этом этапе программа не создаёт окна, но регистрирует класс. Если регистрация прошла успешно, программа завершится без ошибок. Если нет — увидим сообщение в консоли.

## Полный код после Шага 3

Теперь твой `main.go` должен выглядеть так (все импорты и объявления):

```go
package main

import (
    "log"
    "runtime"
    "syscall"
    "unsafe"

    "golang.org/x/sys/windows"
)

const (
    WM_DESTROY = 0x0002
    WM_QUIT    = 0x0012
)

var user32 = windows.NewLazySystemDLL("user32.dll")

var (
    procRegisterClassExW = user32.NewProc("RegisterClassExW")
    procDefWindowProcW   = user32.NewProc("DefWindowProcW")
    procPostQuitMessage  = user32.NewProc("PostQuitMessage")
)

//export WndProc
func WndProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
    switch msg {
    case WM_DESTROY:
        procPostQuitMessage.Call(0)
        return 0
    default:
        ret, _, _ := procDefWindowProcW.Call(hWnd, uintptr(msg), wParam, lParam)
        return ret
    }
}

func registerWindowClass() uint16 {
    className, err := windows.UTF16PtrFromString("MyGoWindowClass")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    wc := windows.WNDCLASSEXW{
        Size:        uint32(unsafe.Sizeof(windows.WNDCLASSEX{})),
        Style:       0,
        WndProc:     syscall.NewCallback(WndProc),
        ClsExtra:    0,
        WndExtra:    0,
        Instance:    0,
        Icon:        0,
        Cursor:      0,
        Background:  0,
        MenuName:    nil,
        ClassName:   className,
        IconSm:      0,
    }

    ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
    if ret == 0 {
        log.Fatal("RegisterClassExW failed")
    }
    return uint16(ret)
}

func main() {
    runtime.LockOSThread()
    registerWindowClass()
}
```

Как Windows узнаёт, какое окно закрывать?

— Windows передаёт hWnd в WndProc. Мы не обрабатываем само закрытие окна, мы лишь реагируем на сообщение WM_DESTROY. Само окно будет уничтожено системой после выхода из цикла сообщений. Пока у нас одно окно, вопрос не стоит, но если бы было несколько, мы бы использовали hWnd для идентификации.

# Шаг 4

**Создание окна.** На этом этапе мы используем зарегистрированный класс для создания видимого окна с заголовком, размерами и стандартными элементами управления (кнопка закрытия, сворачивания, рамка для изменения размера). Также запускается цикл обработки сообщений, который обеспечивает реакцию окна на действия пользователя и системные события.

Теперь, когда класс зарегистрирован, мы можем создать само окно и заставить его отображаться. Для этого нам понадобятся:

1. Константа `WS_OVERLAPPEDWINDOW` — стиль окна (с рамкой, кнопками и т.д.).

2. Константа `CW_USEDEFAULT` — указание Windows самой выбрать позицию и размер.

3. Функция `CreateWindowExW` для создания окна.

4. Функция `ShowWindow` для отображения окна.

5. Цикл сообщений (`GetMessage`, `TranslateMessage`, `DispatchMessage`), который будет обрабатывать события.

Начнём по порядку.

## Часть 1. Добавляем недостающие константы и функции

Открой свой `main.go` и добавь в блок констант:

```go
const (
    WS_OVERLAPPEDWINDOW = 0x00CF0000
    CW_USEDEFAULT       = 0x80000000
    // ... остальные константы уже есть
)
```

В блоке `var` с указателями на функции добавь новые переменные:

```go
var (
    // ... уже объявленные
    procCreateWindowExW  = user32.NewProc("CreateWindowExW")
    procGetMessageW      = user32.NewProc("GetMessageW")
    procTranslateMessage = user32.NewProc("TranslateMessage")
    procDispatchMessageW = user32.NewProc("DispatchMessageW")
)
```

Теперь у нас есть все необходимые функции Windows.

## Часть 2. Функция создания окна

Напишем функцию createWindow, которая создаёт окно и возвращает его идентификатор (`HWND`).

**Объявление:**

```go
func createWindow(className *uint16) uintptr {
    // тело
    return 0
}
```

Заполняем тело:

### 2.1. Заголовок окна

Преобразуем строку `"Тест"` в UTF-16:

```go
title, err := windows.UTF16PtrFromString("Тест")
if err != nil {
    log.Fatal("UTF16PtrFromString failed:", err)
}
```

### 2.2. Вызов CreateWindowExW

Теперь вызываем функцию Windows, передавая параметры:

```go
hwnd, _, _ := procCreateWindowExW.Call(
    0,                               // расширенные стили (0 = нет)
    uintptr(unsafe.Pointer(className)), // имя класса
    uintptr(unsafe.Pointer(title)),      // заголовок окна
    WS_OVERLAPPEDWINDOW,             // стиль окна
    CW_USEDEFAULT, CW_USEDEFAULT,    // x, y (пусть Windows выберет)
    800, 600,                        // ширина, высота
    0, 0, 0, 0,                      // parent, menu, instance, param
)
```

**Пояснение каждого параметра:**

| Параметр | Значение | Объяснение |
|---|---|---|
| 1 | `0` | Расширенные стили (например, для прозрачности, всегда наверху). Нам не нужны. |
| 2 | `uintptr(unsafe.Pointer(className))` | Имя зарегистрированного класса. Указываем, по какому шаблону создавать окно. |
| 3 | `uintptr(unsafe.Pointer(title))` | Текст в заголовке окна. |
| 4 | `WS_OVERLAPPEDWINDOW` | Базовый стиль: рамка, кнопки закрытия/сворачивания, изменяемый размер. |
| 5, 6 | `CW_USEDEFAULT, CW_USEDEFAULT` | Позиция верхнего левого угла. `CW_USEDEFAULT` — Windows сама выбирает. |
| 7, 8 | `800, 600` | Ширина и высота окна в пикселях. |
| 9-12 | `0, 0, 0, 0` | Родительское окно (0 = нет), меню (0 = нет), дескриптор экземпляра (0 = текущий), дополнительные данные (0 = нет). |

### 2.3. Проверка ошибок

Если `hwnd == 0`, создание не удалось:

```go
if hwnd == 0 {
    log.Fatal("CreateWindowExW failed")
}
```

### 2.4. Возврат идентификатора

```go
return hwnd
```

**Итоговый код функции:**

```go
func createWindow(className *uint16) uintptr {
    title, err := windows.UTF16PtrFromString("Тест")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    hwnd, _, _ := procCreateWindowExW.Call(
        0,
        uintptr(unsafe.Pointer(className)),
        uintptr(unsafe.Pointer(title)),
        WS_OVERLAPPEDWINDOW,
        CW_USEDEFAULT, CW_USEDEFAULT,
        800, 600,
        0, 0, 0, 0,
    )

    if hwnd == 0 {
        log.Fatal("CreateWindowExW failed")
    }
    return hwnd
}
```

## Часть 3. Отображение окна

Окно создаётся по умолчанию скрытым. Чтобы оно появилось, вызываем `ShowWindow` с параметром `SW_SHOW`. В пакете `windows` уже есть эта константа.

Добавим в `main()` после создания окна:

```go
windows.ShowWindow(windows.HWND(hwnd), windows.SW_SHOW)
```

**Пояснение:** `windows.HWND(hwnd)` — преобразование `uintptr` в тип `HWND`, ожидаемый функцией `ShowWindow`. `windows.SW_SHOW` — константа, которая говорит «покажи окно в его обычном состоянии».

## Часть 4. Цикл обработки сообщений

После того как окно создано и показано, программа должна войти в бесконечный цикл, который будет получать и обрабатывать сообщения от Windows. Без этого окно не будет реагировать на клики, перерисовку, закрытие.

**Объявление функции:**

```go
func messageLoop() {
    // тело
}
```

### 4.1. Структура для сообщения

Windows использует структуру `MSG` для хранения сообщения. В пакете `windows` она уже есть. Создаём переменную:

```go
var msg windows.Msg
```

### 4.2. Бесконечный цикл

Запускаем цикл, который вызывает GetMessageW, ждёт сообщение и обрабатывает его:

```go
for {
    ret, _, _ := procGetMessageW.Call(
        uintptr(unsafe.Pointer(&msg)), // указатель на структуру
        0,  // все окна
        0, 0, // любые сообщения
    )
    if ret == 0 {
        break // пришёл WM_QUIT - выходим
    }
    procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
    procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
}
```

**Пояснение:**

- `GetMessageW` забирает сообщение из очереди. Если в очереди нет сообщений, он блокирует выполнение (ждёт). Это позволяет не нагружать процессор.

- `ret == 0` означает, что получено сообщение `WM_QUIT` — это сигнал завершить программу. Выходим из цикла.

- `TranslateMessage` преобразует нажатия клавиш в символы (понадобится позже для ввода текста).

- `DispatchMessageW` отправляет сообщение в нашу оконную процедуру `WndProc`.

Итоговый код `messageLoop`:

```go
func messageLoop() {
    var msg windows.Msg
    for {
        ret, _, _ := procGetMessageW.Call(
            uintptr(unsafe.Pointer(&msg)),
            0, 0, 0,
        )
        if ret == 0 {
            break
        }
        procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
        procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
    }
}
```

## Часть 5. Собираем всё в main

Теперь функция `main()` должна:

1. Заблокировать поток (runtime.LockOSThread()).

2. Зарегистрировать класс (registerWindowClass()).

3. Получить имя класса в формате UTF-16.

4. Создать окно (createWindow).

5. Показать окно (ShowWindow).

6. Запустить цикл сообщений (messageLoop).

Код `main`:
```go
func main() {
    runtime.LockOSThread()

    registerWindowClass()

    className, err := windows.UTF16PtrFromString("MyGoWindowClass")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    hwnd := createWindow(className)
    windows.ShowWindow(windows.HWND(hwnd), windows.SW_SHOW)

    messageLoop()
}
```

## Полный код после Шага 4

Вот как должен выглядеть твой `main.go` на данный момент:

```go
package main

import (
    "log"
    "runtime"
    "syscall"
    "unsafe"

    "golang.org/x/sys/windows"
)

const (
    WS_OVERLAPPEDWINDOW = 0x00CF0000
    CW_USEDEFAULT       = 0x80000000
    WM_DESTROY          = 0x0002
    WM_QUIT             = 0x0012
)

var user32 = windows.NewLazySystemDLL("user32.dll")

var (
    procRegisterClassExW = user32.NewProc("RegisterClassExW")
    procCreateWindowExW  = user32.NewProc("CreateWindowExW")
    procDefWindowProcW   = user32.NewProc("DefWindowProcW")
    procGetMessageW      = user32.NewProc("GetMessageW")
    procTranslateMessage = user32.NewProc("TranslateMessage")
    procDispatchMessageW = user32.NewProc("DispatchMessageW")
    procPostQuitMessage  = user32.NewProc("PostQuitMessage")
)

//export WndProc
func WndProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
    switch msg {
    case WM_DESTROY:
        procPostQuitMessage.Call(0)
        return 0
    default:
        ret, _, _ := procDefWindowProcW.Call(hWnd, uintptr(msg), wParam, lParam)
        return ret
    }
}

func registerWindowClass() uint16 {
    className, err := windows.UTF16PtrFromString("MyGoWindowClass")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    wc := windows.WNDCLASSEXW{
        Size:        uint32(unsafe.Sizeof(windows.WNDCLASSEX{})),
        Style:       0,
        WndProc:     syscall.NewCallback(WndProc),
        ClsExtra:    0,
        WndExtra:    0,
        Instance:    0,
        Icon:        0,
        Cursor:      0,
        Background:  0,
        MenuName:    nil,
        ClassName:   className,
        IconSm:      0,
    }

    ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
    if ret == 0 {
        log.Fatal("RegisterClassExW failed")
    }
    return uint16(ret)
}

func createWindow(className *uint16) uintptr {
    title, err := windows.UTF16PtrFromString("Тест")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    hwnd, _, _ := procCreateWindowExW.Call(
        0,
        uintptr(unsafe.Pointer(className)),
        uintptr(unsafe.Pointer(title)),
        WS_OVERLAPPEDWINDOW,
        CW_USEDEFAULT, CW_USEDEFAULT,
        800, 600,
        0, 0, 0, 0,
    )

    if hwnd == 0 {
        log.Fatal("CreateWindowExW failed")
    }
    return hwnd
}

func messageLoop() {
    var msg windows.Msg
    for {
        ret, _, _ := procGetMessageW.Call(
            uintptr(unsafe.Pointer(&msg)),
            0, 0, 0,
        )
        if ret == 0 {
            break
        }
        procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
        procDispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
    }
}

func main() {
    runtime.LockOSThread()

    registerWindowClass()

    className, err := windows.UTF16PtrFromString("MyGoWindowClass")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    hwnd := createWindow(className)
    windows.ShowWindow(windows.HWND(hwnd), windows.SW_SHOW)

    messageLoop()
}
```