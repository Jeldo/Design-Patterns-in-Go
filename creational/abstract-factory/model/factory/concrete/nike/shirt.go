package nike

import (
	"DesignPatterns/creational/abstract-factory/model/factory/shirt"
)

type NikeShirt struct {
	*shirt.Shirt
}
