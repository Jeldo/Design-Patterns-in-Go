package nike

import (
	"DesignPatterns/structural/abstract-factory/model/factory/shoe"
)

type NikeShoe struct {
	*shoe.Shoe
}
