package main

import "fmt"

func plus(a int, b int) (int, int) {
	return a + b, 88
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	res, _ := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
