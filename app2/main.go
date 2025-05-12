package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2

	userHeigt := 1.8
	var userKg float64 = 100

	IMT := userKg / math.Pow(userHeigt, 2)
	fmt.Print(IMT)
}
