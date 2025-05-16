package main

import (
	"encoding/json"
	"fmt"

	"goenumtest/internal/coffee"
)

type Drink struct {
	Type coffee.Coffee `json:"type"`
}

func PrintCoffee(c coffee.Coffee) {
	fmt.Println("Coffee type is:", c)
}

func main() {
	PrintCoffee(coffee.Drip)
	PrintCoffee(coffee.Latte)

	payload := []byte(`{"type":"cappuccino"}`)

	drink := Drink{}
	if err := json.Unmarshal(payload, &drink); err != nil {
		panic(err)
	}

	fmt.Println("Unmarshaled drink is:", drink)

	s, err := json.Marshal(drink)
	if err != nil {
		panic(err)
	}

	fmt.Println("Marshaled drink is:", string(s))
}
