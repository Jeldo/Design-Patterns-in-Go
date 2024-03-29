package main

import (
	"DesignPatterns/creational/builder/house"
	"fmt"
)

func main() {
	normalBuilder := house.GetBuilder("normal")
	iglooBuilder := house.GetBuilder("igloo")

	director := house.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Println("Normal House Door Type:", normalHouse.DoorType())
	fmt.Println("Normal House Window Type:", normalHouse.WindowType())
	fmt.Println("Normal House Num Floor:", normalHouse.Floor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Println("Igloo House Door Type:", iglooHouse.DoorType())
	fmt.Println("Igloo House Window Type:", iglooHouse.WindowType())
	fmt.Println("Igloo House Num Floor:", iglooHouse.Floor())
}
