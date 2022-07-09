package concretes

import (
	"DesignPatterns/structural/factory-method/model"
)

type AK47 struct {
	model.Gun
}

func NewAK47() model.IGun {
	return &AK47{Gun: *model.NewGun("AK47", 3)}
}
