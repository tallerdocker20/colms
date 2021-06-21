package main

import "fmt"

func main() {
	var numero int
	var puntero *int // Puntero a un int (en C int *x;)

	numero = 10
	puntero = &numero

	fmt.Println("numero:", numero)     // 10
	fmt.Println("&numero:", &numero)   // 0xc0000a6058
	fmt.Println("&puntero:", &puntero) // 0xc0000d0018
	fmt.Println("puntero:", puntero)   // 0xc0000a6058
	fmt.Println("*puntero:", *puntero) // 10
	fmt.Println("------")
	*puntero = 20
	fmt.Println("*puntero:", *puntero) // 20
	fmt.Println("numero:", numero)     // 20

}
