package main

import (
    "log"       // для вывода ошибок
    "runtime"   // для LockOSThread
    "syscall"   // для преобразования Go-функции в указатель для Windows
    "unsafe"    // для работы с сырой памятью

    "golang.org/x/sys/windows" // готовые обёртки для WinAPI
)

// Константы, необходимые для работы с окнами
const (
    WS_OVERLAPPEDWINDOW = 0x00CF0000
    CW_USEDEFAULT       = 0x80000000
    WM_DESTROY          = 0x0002
    WM_QUIT             = 0x0012
)

// Загружаем библиотеку user32.dll
var user32 = windows.NewLazySystemDLL("user32.dll")

// Указатели на функции из user32.dll
var (
    procRegisterClassExW = user32.NewProc("RegisterClassExW")
    procCreateWindowExW  = user32.NewProc("CreateWindowExW")
    procDefWindowProcW   = user32.NewProc("DefWindowProcW")
    procGetMessageW      = user32.NewProc("GetMessageW")
    procTranslateMessage = user32.NewProc("TranslateMessage")
    procDispatchMessageW = user32.NewProc("DispatchMessageW")
    procPostQuitMessage  = user32.NewProc("PostQuitMessage")
    
)

// Оконная процедура — обработчик событий окна
//export WndProc
func WndProc(hWnd uintptr, msg uint32, wParam uintptr, lParam uintptr) uintptr {
    switch msg {
    case WM_DESTROY:
        // при закрытии окна отправляем команду выхода
        procPostQuitMessage.Call(0)
        return 0
    default:
        // остальные сообщения передаём стандартному обработчику
        ret, _, _ := procDefWindowProcW.Call(hWnd, uintptr(msg), wParam, lParam)
        return ret
    }
}

// Регистрация класса окна
func registerWindowClass() uint16 {
    // Преобразуем строку в UTF-16 (Windows ожидает такой формат)
    className, err := windows.UTF16PtrFromString("MyGoWindowClass")
    if err != nil {
        log.Fatal("UTF16PtrFromString failed:", err)
    }

    // Структура WNDCLASSEXW — описание класса
    wc := windows.WNDCLASSEXW{
        Size:        uint32(unsafe.Sizeof(windows.WNDCLASSEX{})),
        Style:       0,
        WndProc:     syscall.NewCallback(WndProc), // указатель на обработчик
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

    // Вызов RegisterClassExW
    ret, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
    if ret == 0 {
        log.Fatal("RegisterClassExW failed")
    }
    return uint16(ret)
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

