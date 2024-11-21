package main

import (
	"fmt"

	"github.com/ewollesen/building-with-private-modules/caution"
)

func main() {
	fmt.Println("Hello, World!")

	var thinger caution.Thinger = caution.NewThinger()
	fmt.Println("the thing is,", thinger.Thing())
}
