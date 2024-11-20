//go:build private

// !!! THIS FILE MUST BE TAGGED WITH go:build private !!!
//
// This file integrates with a Tidepool partner that is unwilling to allow us to publish our
// integration with their SDK. See doc.go for further conventions and suggestions.
//
// This file, despite the "private" build tag is still publicly visible code.
//
// In this example, "github.com/ewollesen/optional-private-module" is a private Go
// package. This package whether it's ours or our partner's, contains code we're
// contractually obligated not to keep private.

package caution

import (
	"fmt"

	"github.com/ewollesen/optional-private-module/secret"
)

func UsePrivateCode() {
	fmt.Println(secret.Hello())
}

// PartnerThinger is a placeholder implementation of Thinger.
type PartnerThinger struct{}

// Thing implements Thinger.
func (t *PartnerThinger) Thing() string {
	return "This calls the real partner thinger."
}

// NewThinger instantiates a placholder thinger.
func NewThinger() *PartnerThinger {
	return &PartnerThinger{}
}
