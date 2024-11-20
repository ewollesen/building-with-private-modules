package main

import (
	"fmt"

	"github.com/ewollesen/building-with-private-modules/caution"
)

func main() {
	fmt.Println("Hello, World!")

	// Let's call the secret method. If we have the secret module, this will include the
	// secret. If not, then it should use a placeholder.
	caution.UsePrivateCode()

	var thinger caution.Thinger = caution.NewThinger()
	fmt.Println("the thing is, ", thinger.Thing())
}
