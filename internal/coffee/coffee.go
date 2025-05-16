package coffee

import (
	"fmt"
	"strings"
)

//go:generate stringer -linecomment -type=Coffee

type Coffee int

const (
	Drip       Coffee = iota // drip coffee
	Latte                    // latte
	Breve                    // breve
	Cappuccino               // cappuccino
)

func (c Coffee) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Coffee) UnmarshalText(text []byte) error {
	want := string(text)
	for i := 0; i < len(_Coffee_index)-1; i++ {
		name := _Coffee_name[_Coffee_index[i]:_Coffee_index[i+1]]
		if strings.EqualFold(name, want) {
			*c = Coffee(i)
			return nil
		}
	}
	return fmt.Errorf("invalid Coffee %q", want)
}
