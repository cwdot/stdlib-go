package main

import (
	"github.com/cwdot/go-stdlib/wood"
)

func main() {
	wood.Init(wood.InfoLevel)

	wood.Printf("test %s", "f")
	wood.Println("test", "f")
}