package nike

import (
	"DesignPatterns/creational/abstract-factory/model/factory/shirt"
	"DesignPatterns/creational/abstract-factory/model/factory/shoe"
	"DesignPatterns/creational/abstract-factory/model/factory/sports"
)

type NikeFactory struct{}

func (f *NikeFactory) CreateShoe() sports.Shoe {
	return &NikeShoe{
		Shoe: shoe.NewShoe("NIKE SHOE LOGO", 22),
	}
}

func (f *NikeFactory) CreateShirt() sports.Shirt {
	return &NikeShirt{
		Shirt: shirt.NewShirt("NIKE SHIRT LOGO", 222),
	}
}
