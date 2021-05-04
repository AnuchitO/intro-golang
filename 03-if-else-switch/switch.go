package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday

	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Printf("สวัสดีวัน %v อยู่ดีมีแฮงเดย์\n", today)
	}
}
