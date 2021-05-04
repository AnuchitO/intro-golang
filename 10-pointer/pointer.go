package main

import "fmt"

func main() {
	me := "Gopher"
	fmt.Printf("You are %s\n", me)

	addr := &me
	fmt.Println(addr)
	fmt.Printf("%T\n", addr)

	*addr = "Penguin"
	fmt.Printf("You are %s\n", me)
}
