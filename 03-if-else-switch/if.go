package main

import "fmt"

func main() {
	num := 33

	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num == 31 {
		fmt.Println("คนมีคู่ไม่รู้หรอก")
	} else {
		fmt.Println("เลขโสด")
	}

}
