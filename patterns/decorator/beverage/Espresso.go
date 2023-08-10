package beverage

type Espresso struct {
}

func GetNewEspresso() *Espresso {
	return &Espresso{}
}

func (espresso *Espresso) Cost() int {
	return 45
}
