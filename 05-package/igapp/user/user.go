package user

import "fmt"

func Info(name string, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Printf("I'm %d year old\n", age)
}
