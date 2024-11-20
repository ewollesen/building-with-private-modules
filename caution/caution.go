// Package caution will optionally link to private modules.
//
// Building with the "private" tag will link the private Go module into the final
// binary. Building without the "private" tag, will leave out the private code, instead
// substituting a placeholder function.
//
// Any and all references to the private module should be made from this file, or files
// similarly guarded with the "//go:build private" tag.
//
// Any symbols exported by this package, should have two versions: one that is guarded with
// the "private" tag and references the private Go module, and another that is guarded with
// the "!private" tag and uses placeholders.

//go:build private

package caution

import (
	"fmt"

	"github.com/ewollesen/optional-private-module/secret"
)

func UsePrivateCode() {
	fmt.Println(secret.Hello())
}
