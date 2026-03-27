package main

import (
    "fmt"
	"errors"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

func main() {
	answer, err := divide(1,0)
	fmt.Println(answer, err)
}

type error interface {
    Error() string
}
