package main

import (
	"fmt"

	"goenumtest/internal/coffee"
)

func PrintCoffee(c coffee.Coffee) {
	fmt.Println("Coffee type is:", c)
}

func main() {
	PrintCoffee(coffee.Drip)
	PrintCoffee(coffee.Latte)
}
