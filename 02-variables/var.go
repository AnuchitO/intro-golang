package main

import "fmt"

/*
	The zero value is:
	0 for numeric types,
	false for the boolean type
	"" (the empty string) for strings.
*/

/* basic type
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func main() {
	var name string
	s := "Gopher นำแน่!!!"
	i := 5
	f := 3.7
	b := true
	r := '界'

	fmt.Printf("name: %q\n", name)
	fmt.Printf("type: %T\n", name)

	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("string: %v\n", s)
	fmt.Printf("rune: %v\n", r)

}
