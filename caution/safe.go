//go:build !private

// !!! THIS FILE MUST BE TAGGED WITH go:build !private !!!
//
// This file is the complement of private.go. It contains public placeholder implementations
// of optional integrations to Tidepool partners that are unwilling to allow us to publish
// our integration with their SDKs. See doc.go for further conventions and suggestions.

package caution

import "fmt"

// UsePrivateCode is a placeholder implementation.
func UsePrivateCode() {
	fmt.Println("I'm a stub without private code.")
}

// PlaceholderThinger is a placeholder implementation of Thinger.
type PlaceholderThinger struct{}

// Thing implements Thinger with a placeholder. It is *NOT* the Real Thing™.
func (t *PlaceholderThinger) Thing() string {
	return "placeholder thinger"
}

// NewThinger instantiates a placholder thinger.
func NewThinger() *PlaceholderThinger {
	return &PlaceholderThinger{}
}
