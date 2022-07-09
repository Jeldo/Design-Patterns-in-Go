package factory

import (
	"DesignPatterns/structural/abstract-factory/model/factory/concrete/adidas"
	"DesignPatterns/structural/abstract-factory/model/factory/concrete/nike"
	"DesignPatterns/structural/abstract-factory/model/factory/sports"
	"errors"
)

func CreateSportsFactory(brand string) (sports.SportsFactory, error) {
	if brand == "adidas" {
		return &adidas.AdidasFactory{}, nil
	}
	if brand == "nike" {
		return &nike.NikeFactory{}, nil
	}

	return nil, errors.New("invalid brand type")
}
