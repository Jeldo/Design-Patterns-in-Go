package gun_factory

import (
	concretes2 "DesignPatterns/structural/factory-method/concretes"
	"DesignPatterns/structural/factory-method/model"
	"errors"
)

func GetGun(gunType string) (model.IGun, error) {
	if gunType == "AK47" {
		return concretes2.NewAK47(), nil
	}
	if gunType == "MUSKET" {
		return concretes2.NewMusket(), nil
	}

	return nil, errors.New("invalid gun type")
}
