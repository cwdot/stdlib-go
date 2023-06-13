package main

import (
	"github.com/sirupsen/logrus"

	"github.com/cwdot/go-stdlib/wood"
)

func main() {
	wood.Init(logrus.InfoLevel)

	wood.Printf("test %s", "f")
	wood.Println("test", "f")
}
