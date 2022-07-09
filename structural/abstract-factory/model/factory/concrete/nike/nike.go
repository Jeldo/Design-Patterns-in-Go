package nike

import (
	"DesignPatterns/structural/abstract-factory/model/factory/shirt"
	"DesignPatterns/structural/abstract-factory/model/factory/shoe"
	"DesignPatterns/structural/abstract-factory/model/factory/sports"
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
