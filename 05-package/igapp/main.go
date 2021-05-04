package main

import (
	"fmt"

	"github.com/anuchito/igapp/time"
	"github.com/anuchito/igapp/user"
)

func main() {
	user.Info("Nong", "โกเฟอร์", 15)

	t := time.Today()
	fmt.Println("today is ", t)
}
