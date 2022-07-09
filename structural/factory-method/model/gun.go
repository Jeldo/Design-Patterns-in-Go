package model

type IGun interface {
	SetName(name string)
	SetPower(power int)
	Name() string
	Power() int
}

type Gun struct {
	name  string
	power int
}

func NewGun(name string, power int) *Gun {
	return &Gun{name: name, power: power}
}

func (g *Gun) Name() string {
	return g.name
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) Power() int {
	return g.power
}

func (g *Gun) SetPower(power int) {
	g.power = power
}
