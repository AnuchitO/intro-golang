package main

import "fmt"

func show(name string) {
	fmt.Printf("value: %v\n", name)
}

func cal(op func(int, int) int) {
	r := op(3, 4)
	fmt.Println("result =", r)
}

func main() {
	var name = "Nong"
	fmt.Printf("value: %v\n", name)
	fmt.Printf("type: %T\n", name)
	show(name)
	fmt.Println()

	var plus = func(a int, b int) int {
		return a + b
	}

	p := plus(1, 2)
	fmt.Println("1+2 =", p)
	fmt.Printf("type: %T\n", plus)

	var minus = func(a int, b int) int {
		return a - b
	}
	cal(minus)
}
