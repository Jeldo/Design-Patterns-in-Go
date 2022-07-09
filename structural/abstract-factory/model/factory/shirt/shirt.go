package shirt

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) Logo() string {
	return s.logo
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) Size() int {
	return s.size
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func NewShirt(logo string, size int) *Shirt {
	return &Shirt{logo: logo, size: size}
}
