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
	// iterate all valid values (0 .. len(_Coffee_index)-2)
	for i := Coffee(0); i < Coffee(len(_Coffee_index)-1); i++ {
		name := _Coffee_name[_Coffee_index[i]:_Coffee_index[i+1]]
		if strings.EqualFold(name, want) { // or if name==want for exact match
			*c = i
			return nil
		}
	}
	return fmt.Errorf("invalid Coffee %q", want)
}
