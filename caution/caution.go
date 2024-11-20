//go:build private

// Package caution will optionally link to private modules.

package caution

import (
	"fmt"

	"github.com/ewollesen/optional-private-module/secret"
)

func UsePrivateCode() {
	fmt.Println(secret.Hello())
}
