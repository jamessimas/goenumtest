package coffee

//go:generate stringer -linecomment -type=Coffee

type Coffee int

const (
	Drip       Coffee = iota // drip coffee
	Latte                    // latte
	Breve                    // breve
	Cappuccino               // cappuccino
)
