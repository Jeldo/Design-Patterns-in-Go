package pizza

type TomatoTopping struct {
	Pizza Pizza
}

func (tt *TomatoTopping) GetPrice() int {
	return tt.Pizza.GetPrice() + 7
}
