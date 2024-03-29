package adidas

import (
	"DesignPatterns/creational/abstract-factory/model/factory/shirt"
	"DesignPatterns/creational/abstract-factory/model/factory/shoe"
	"DesignPatterns/creational/abstract-factory/model/factory/sports"
)

type AdidasFactory struct{}

func (a *AdidasFactory) CreateShoe() sports.Shoe {
	return &AdidasShoe{
		Shoe: shoe.NewShoe("ADIDAS SHOE LOGO", 11),
	}
}

func (a *AdidasFactory) CreateShirt() sports.Shirt {
	return &AdidasShirt{
		Shirt: shirt.NewShirt("ADIDAS SHIRT LOGO", 111),
	}
}
