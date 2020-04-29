package mypkg

import (
	"fmt"

	"github.com/leki75/mypkg/v2/math"
)

func Print(s string, i int) {
	fmt.Println(s, i)
}

func Sum(i, j int) int {
	return math.Add(i, j)
}
