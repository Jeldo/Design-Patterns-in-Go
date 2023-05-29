package sports

type SportsFactory interface {
	CreateShoe() Shoe
	CreateShirt() Shirt
}

type Shoe interface {
	SetLogo(logo string)
	SetSize(size int)
	Logo() string
	Size() int
}

type Shirt interface {
	SetLogo(logo string)
	SetSize(size int)
	Logo() string
	Size() int
}
