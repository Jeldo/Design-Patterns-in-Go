package adidas

import (
	"DesignPatterns/creational/abstract-factory/model/factory/shirt"
)

type AdidasShirt struct {
	*shirt.Shirt
}
