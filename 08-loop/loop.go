package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "javascript", "F#"}
	fmt.Println("\nfor basic")
	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Println(i, ":", value)
	}

	fmt.Println("\nfor range slice")
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}

	fmt.Println("\nonly value")
	for _, value := range langs {
		fmt.Println("only value :", value)
	}

	fmt.Println("\nonly index")
	for index := range langs {
		fmt.Println("only index :", index)
	}

}
