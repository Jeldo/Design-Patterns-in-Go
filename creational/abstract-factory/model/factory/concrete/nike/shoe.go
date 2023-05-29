package nike

import (
	"DesignPatterns/creational/abstract-factory/model/factory/shoe"
)

type NikeShoe struct {
	*shoe.Shoe
}
