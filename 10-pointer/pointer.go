package main

import "fmt"

func main() {
	me := "Gopher"
	fmt.Printf("You are %s\n", me)

	// var addr *string = &me
	addr := &me
	fmt.Println(addr)
	fmt.Printf("%T\n", addr)

	// อ่านค่าที่ addr ชึ้อยู่
	v := *addr
	fmt.Println(v)

	// เป็นค่าที่ addr ชึ้อยู่
	*addr = "Penguin"
	fmt.Printf("You are %s\n", me)
}
