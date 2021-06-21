package main

import "fmt"

func main() {
	f := "apple"
	f = "ios"

	var v string
	v = "wind"

	fmt.Println("hello world" + f + v)

	var a [5]int
	a[4] = 100
	fmt.Println("a:", a)
	fmt.Println("get:", a[4])

	s := make([]string, 5)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"
	fmt.Println("emp:", s)

	l := s[:2]
	fmt.Println("sl1:", l)

	n := map[string]string{"foo": "5", "bar": "5"}
	fmt.Println("map:", n)

}
