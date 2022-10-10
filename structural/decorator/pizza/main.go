package main

import (
	"DesignPatterns/structural/decorator/pizza/pizza"
	"fmt"
)

func main() {
	veggiePizza := &pizza.VeggieMania{}
	pizzaWithCheese := &pizza.CheeseTopping{Pizza: veggiePizza}
	pizzaWithCheeseAndTomato := &pizza.TomatoTopping{Pizza: pizzaWithCheese}
	fmt.Printf("total price=%d", pizzaWithCheeseAndTomato.GetPrice())
}
