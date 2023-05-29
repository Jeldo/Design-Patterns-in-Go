package gun_factory

import (
	"DesignPatterns/creational/factory-method/concretes"
	"DesignPatterns/creational/factory-method/model"
	"errors"
)

func GetGun(gunType string) (model.IGun, error) {
	if gunType == "AK47" {
		return concretes.NewAK47(), nil
	}
	if gunType == "MUSKET" {
		return concretes.NewMusket(), nil
	}

	return nil, errors.New("invalid gun type")
}
