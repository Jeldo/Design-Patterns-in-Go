package adidas

import (
	"DesignPatterns/creational/abstract-factory/model/factory/shoe"
)

type AdidasShoe struct {
	*shoe.Shoe
}
