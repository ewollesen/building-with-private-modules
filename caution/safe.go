//go:build !private

package caution

import "fmt"

func UsePrivateCode() {
	fmt.Println("I'm a stub without private code.")
}

// PublicThinger implements Thinger.
type PublicThinger struct{}

func (t *PublicThinger) Thing() string {
	return "placeholder thinger"
}
