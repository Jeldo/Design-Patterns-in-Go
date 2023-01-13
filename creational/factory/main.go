package main

import (
	"errors"
)

type Gun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

type AK47 struct {
	gun
}

func NewAK47() Gun {
	return &AK47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	gun
}

func NewMusket() Gun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (Gun, error) {
	if gunType == "ak47" {
		return NewAK47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, errors.New("wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g Gun) {
	println("Gun:", g.getName())
	println("Power:", g.getPower())
}
