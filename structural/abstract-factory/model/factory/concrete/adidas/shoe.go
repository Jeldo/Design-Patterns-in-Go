package adidas

import (
	"DesignPatterns/structural/abstract-factory/model/factory/shoe"
)

type AdidasShoe struct {
	*shoe.Shoe
}
