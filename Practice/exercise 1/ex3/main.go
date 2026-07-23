package main

import "fmt"

var (
	b byte = 255
	small int32 = 2_147_483_647
	bigI uint64 = 18_446_744_073_709_551_615
)


func main() {

	fmt.Println("b byte = ", b)
	fmt.Println("small int32 = ", small)
	fmt.Println("bigI uint64 = ", bigI)

	b += 1
	small += 1
	bigI += 1

	fmt.Println("b byte + 1 = ", b)
	fmt.Println("small int32 + 1 = ", small)
	fmt.Println("bigI uint64 + 1 = ", bigI)
}