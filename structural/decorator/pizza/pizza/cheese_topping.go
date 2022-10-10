package pizza

type CheeseTopping struct {
	Pizza Pizza
}

func (ct *CheeseTopping) GetPrice() int {
	return ct.Pizza.GetPrice() + 10
}
