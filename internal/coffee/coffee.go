package coffee

//go:generate stringer -type=Coffee

type Coffee int

const (
	Drip Coffee = iota
	Latte
	Breve
	Cappuccino
)
