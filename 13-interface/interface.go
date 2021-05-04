package main

import "fmt"

type Phone interface {
	Call(number string)
}

type Samsung struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Println(s.Name, "calling", number)
}

func (s Samsung) Answer() {
	fmt.Println(s.Name, "hello there!")
}

func Dial(p Phone) {
	p.Call("+6678787879")
}

func main() {
	s := Samsung{
		Name: "S10",
	}

	s.Answer()
	s.Call("sss")

	Dial(s)
}
