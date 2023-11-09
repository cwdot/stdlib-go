package main

import (
	"github.com/cwdot/stdlib-go/wood"
)

func main() {
	wood.Init()

	wood.Printf("test %s", "f")
	wood.Println("test", "f")
}
