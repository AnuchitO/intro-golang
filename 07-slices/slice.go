package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "javascript"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[0]
	fmt.Printf("langs[0]: %#v\n", n)

	langs[1] = "Pythonistas"
	fmt.Printf("langs: %#v\n", langs)

	l := len(langs)
	fmt.Println("length of langs :", l)

	langs = append(langs, "F#", "Em", "Am")
	fmt.Printf("langs: %#v\n", langs)

	fmt.Println("length of langs :", len(langs))

}
