package main

import "fmt"

func showAll(ns [4]string) {
	fmt.Printf("show all : %#v\n", ns)
}

func main() {
	var langs = [4]string{"golang", "python", "javascript"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[1]
	fmt.Printf("%#v\n", n)

	langs[1] = "Pythonistas"
	fmt.Printf("%#v\n", langs)

	showAll(langs)

	cords := [4]string{"Am", "Gm", "F7-11"}
	showAll(cords)
}
