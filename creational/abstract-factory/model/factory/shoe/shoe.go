package shoe

type Shoe struct {
	logo string
	size int
}

func (c *Shoe) Logo() string {
	return c.logo
}

func (c *Shoe) SetLogo(logo string) {
	c.logo = logo
}

func (c *Shoe) Size() int {
	return c.size
}

func (c *Shoe) SetSize(size int) {
	c.size = size
}

func NewShoe(logo string, size int) *Shoe {
	return &Shoe{logo: logo, size: size}
}
