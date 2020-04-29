package main

import (
	"github.com/leki75/mypkg/v2"
	"github.com/leki75/mypkg/v2/math"
)

func main() {
	mypkg.Print("Math", math.Add(1, 2))
}
