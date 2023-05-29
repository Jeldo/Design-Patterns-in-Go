package concretes

import (
	"DesignPatterns/creational/factory-method/model"
)

type Musket struct {
	model.Gun
}

func NewMusket() model.IGun {
	return &Musket{Gun: *model.NewGun("musket", 1)}
}
