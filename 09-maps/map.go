package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}
	fmt.Printf("% #v\n", status)

	l := len(status)
	fmt.Printf("length: %d\n\n", l)

	status[200] = "Okie"
	status[285] = "เมาแน่เลย"
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	key := 285
	delete(status, key)
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	// var m map[string]string = make(map[string]string) // จองที่ mem ให้แล้ว ไม่ nil แล้ว
	// var m = make(map[string]string)
	// m := make(map[string]string)
	// m := map[string]string{}
	m := map[string]string{"ออเจ้า": "+66-978788997", "Steven": "+61-857985555"}

	if m == nil {
		fmt.Printf("m is nil. value : %#v\n", m)
	} else {
		fmt.Printf("m is not nil. value : %#v\n", m)
	}

	for key, value := range status {
		fmt.Printf("key: %d, value: %s\n", key, value)
	}
	fmt.Println()

	// omit key
	for _, value := range status {
		fmt.Printf("only value: %s\n", value)
	}
	fmt.Println()

	// omit value
	for key := range status {
		fmt.Printf("only key: %d\n", key)
	}

}
