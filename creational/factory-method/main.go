package main

import (
	"DesignPatterns/creational/factory-method/gun-factory"
	"DesignPatterns/creational/factory-method/model"
	"fmt"
)

func main() {
	ak47, _ := gun_factory.GetGun("AK47")
	musket, _ := gun_factory.GetGun("MUSKET")

	printDetails(ak47)
	printDetails(musket)
}
func printDetails(g model.IGun) {
	fmt.Printf("Gun: %s", g.Name())
	fmt.Println()
	fmt.Printf("Power: %d", g.Power())
	fmt.Println()
}
