package main

import (
	"DesignPatterns/structural/abstract-factory/model/factory"
	"DesignPatterns/structural/abstract-factory/model/factory/sports"
	"fmt"
)

func main() {
	adidasFactory, _ := factory.CreateSportsFactory("adidas")
	adidasShoe := adidasFactory.CreateShoe()
	adidasShirt := adidasFactory.CreateShirt()
	printDetails(adidasShoe, adidasShirt)

	nikeFactory, _ := factory.CreateSportsFactory("nike")
	nikeShoe := nikeFactory.CreateShoe()
	nikeShirt := nikeFactory.CreateShirt()
	printDetails(nikeShoe, nikeShirt)

	_, err := factory.CreateSportsFactory("invalid")
	if err != nil {
		fmt.Println(err)
	}
}

func printDetails(shoe sports.Shoe, shirt sports.Shirt) {
	fmt.Printf("logo: %s, size: %v\n", shoe.Logo(), shoe.Size())
	fmt.Printf("logo: %s, size: %v\n", shirt.Logo(), shirt.Size())
}
