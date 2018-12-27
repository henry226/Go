package main

import (
	"fmt"
	"math"

	"github.com/henry226/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println("math.Floor(2.7) = ", math.Floor(2.7))
	fmt.Println("math.Ceil(2.7) = ", math.Ceil(2.7))
	fmt.Println("math.Sqrt(16) = ", math.Sqrt(16))
	fmt.Println("Reverse 'hello' = ", strutil.Reverse("hello"))
}
